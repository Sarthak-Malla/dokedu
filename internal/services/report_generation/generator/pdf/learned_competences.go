package pdf

import (
	"context"
	"example/internal/db"
	"fmt"
	"github.com/uptrace/bun"
)

func (g *Generator) LearnedCompetencesReportData(report db.Report) (*CompetencesTemplateData, error) {
	ctx := context.Background()

	var data CompetencesTemplateData

	var student db.User
	err := g.cfg.DB.NewSelect().Model(&student).Where("id = ?", report.StudentUserID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	data.StudentName = fmt.Sprintf("%s %s", student.FirstName, student.LastName)

	var userStudent db.UserStudent
	err = g.cfg.DB.NewSelect().Model(&userStudent).Where("user_id = ?", report.StudentUserID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	data.StudentName = fmt.Sprintf("%s %s", student.FirstName, student.LastName)
	data.StudentBirthday = userStudent.Birthday.Format("02.01.2006")

	// a school year is from 1st of August to 31st of July and is 2006/06

	date := report.CreatedAt
	if date.Month() < 8 {
		date = date.AddDate(-1, 0, 0)
	}
	lastTwoDigitsOfNextYear := (date.Year() + 1) % 100

	data.SchoolYear = fmt.Sprintf("%d/%d", date.Year(), lastTwoDigitsOfNextYear)

	var organisation db.Organisation
	err = g.cfg.DB.NewSelect().Model(&organisation).Where("id = ?", report.OrganisationID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	data.OrganisationName = organisation.Name
	data.OrganisationAddress = organisation.Address
	data.OrganisationLogoURL = organisation.LogoURL

	var userCompetences []db.UserCompetence
	err = g.cfg.DB.NewSelect().
		Model(&userCompetences).
		Where("user_id = ?", report.StudentUserID).
		Where("organisation_id = ?", report.OrganisationID).
		Where("created_at >= ?", report.From).
		Where("created_at <= (DATE ? + 1)", report.To).
		Order("created_at DESC").
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	// create a user competences ids list with the unique values of userCompetences.CompetenceID
	var userCompetenceIds []string
Outer:
	for _, uc := range userCompetences {
		// check if competence already exists
		for _, id := range userCompetenceIds {
			if id == uc.CompetenceID {
				continue Outer
			}
		}
		userCompetenceIds = append(userCompetenceIds, uc.CompetenceID)
	}

	var competences []db.Competence
	err = g.cfg.DB.NewSelect().
		Model(&competences).
		Where("organisation_id = ?", report.OrganisationID).
		Where("id IN (?)", bun.In(userCompetenceIds)).
		Order("sort_order").Scan(ctx)
	if err != nil {
		return nil, err
	}

	// for each competence, get the parent competence
	// recursively get all parent competences until the parent competence is nil
	// append the parent competence to the competences slice if it doesn't exist yet
	for i := range competences {
		parents, err := g.getParentCompetences(report, competences[i])
		if err != nil {
			return nil, err
		}

		for _, p := range parents {
			// check if parent already exists
			var exists bool
			for _, c := range competences {
				if c.ID == p.ID {
					exists = true
					break
				}
			}
			if !exists {
				competences = append(competences, p)
			}
		}
	}

	// create a map with key competence.ParentID and value []db.Competence
	var competencesMap = make(map[string][]db.Competence)
	for _, c := range competences {
		competencesMap[c.CompetenceID.String] = append(competencesMap[c.CompetenceID.String], c)
	}

	// create a map with key userCompetence.CompetenceID and value []db.UserCompetence
	var userCompetencesMap = make(map[string][]db.UserCompetence)
	for _, uc := range userCompetences {
		userCompetencesMap[uc.CompetenceID] = append(userCompetencesMap[uc.CompetenceID], uc)
	}

	subjects := Subjects(competences)

	for i := range subjects {
		subject := subjects[i]

		if data.Competences == nil {
			data.Competences = make([]Competence, len(subjects))
		}
		data.Competences[i].Name = subject.Name

		data.Competences[i].Color = subject.Color.String

		if !subject.Color.Valid {
			data.Competences[i].Color = "stone"
		}

		competences, err := g.populateCompetences(report, subject, competencesMap, userCompetencesMap, subject.Color.String)
		if err != nil {
			return nil, err
		}

		if data.Competences[i].Competences == nil {
			data.Competences[i].Competences = make([]Competence, len(competences))
		}
		data.Competences[i].Competences = competences
	}

	return &data, nil
}

func (g *Generator) getParentCompetences(report db.Report, competence db.Competence) ([]db.Competence, error) {
	query := `SELECT * FROM get_competence_tree_reverse(?);`

	var parents []db.Competence
	err := g.cfg.DB.NewRaw(query, competence.ID).Scan(context.Background(), &parents)
	if err != nil {
		return nil, err
	}

	return parents, nil
}
