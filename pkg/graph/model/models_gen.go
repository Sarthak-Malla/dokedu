// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"example/pkg/db"
	"time"
)

type ChatConnection struct {
	Edges      []*db.Chat `json:"edges,omitempty"`
	PageInfo   *PageInfo  `json:"pageInfo"`
	TotalCount int        `json:"totalCount"`
}

type CompetenceConnection struct {
	Edges      []*db.Competence `json:"edges,omitempty"`
	PageInfo   *PageInfo        `json:"pageInfo"`
	TotalCount int              `json:"totalCount"`
}

type CompetenceFilterInput struct {
	Type    []*db.CompetenceType `json:"type,omitempty"`
	Parents []*string            `json:"parents,omitempty"`
}

type CreateEntryFileInput struct {
	File string `json:"file"`
}

type CreateEntryInput struct {
	Date            time.Time                         `json:"date"`
	Body            *string                           `json:"body,omitempty"`
	User            string                            `json:"user"`
	Users           []string                          `json:"users"`
	Tags            []string                          `json:"tags"`
	Files           []*CreateEntryFileInput           `json:"files"`
	UserCompetences []*CreateEntryUserCompetenceInput `json:"userCompetences"`
}

type CreateEntryUserCompetenceInput struct {
	Level      int    `json:"level"`
	User       string `json:"user"`
	Competence string `json:"competence"`
}

type CreateReportInput struct {
	Format      db.ReportFormat `json:"format"`
	Kind        db.ReportKind   `json:"kind"`
	From        time.Time       `json:"from"`
	To          time.Time       `json:"to"`
	Meta        string          `json:"meta"`
	FilterTags  []string        `json:"filterTags"`
	StudentUser string          `json:"studentUser"`
}

type CreateUserInput struct {
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	Email     string      `json:"email"`
	Role      db.UserRole `json:"role"`
	Birthday  *time.Time  `json:"birthday,omitempty"`
	LeftAt    *time.Time  `json:"leftAt,omitempty"`
	JoinedAt  *time.Time  `json:"joinedAt,omitempty"`
}

type EntryConnection struct {
	Edges      []*db.Entry `json:"edges,omitempty"`
	PageInfo   *PageInfo   `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

type EntryFilterInput struct {
	Author      []*string  `json:"author,omitempty"`
	Users       []*string  `json:"users,omitempty"`
	Tags        []*string  `json:"tags,omitempty"`
	Competences []*string  `json:"competences,omitempty"`
	From        *time.Time `json:"from,omitempty"`
	To          *time.Time `json:"to,omitempty"`
}

type EventConnection struct {
	Edges      []*db.Event `json:"edges,omitempty"`
	PageInfo   *PageInfo   `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

type EventFilterInput struct {
	From        *time.Time `json:"from,omitempty"`
	To          *time.Time `json:"to,omitempty"`
	ShowDeleted *bool      `json:"showDeleted,omitempty"`
}

type OrganisationConnection struct {
	Edges      []*db.Organisation `json:"edges,omitempty"`
	PageInfo   *PageInfo          `json:"pageInfo"`
	TotalCount int                `json:"totalCount"`
}

type PageInfo struct {
	HasNextPage     bool `json:"hasNextPage"`
	HasPreviousPage bool `json:"hasPreviousPage"`
	CurrentPage     int  `json:"currentPage"`
}

type ReportConnection struct {
	Edges      []*db.Report `json:"edges,omitempty"`
	PageInfo   *PageInfo    `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInPayload struct {
	Token string `json:"token"`
}

type SignUpInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateEntryInput struct {
	Date *time.Time `json:"date,omitempty"`
	Body *string    `json:"body,omitempty"`
	User *string    `json:"user,omitempty"`
}

type UpdateUserInput struct {
	ID        string     `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     *string    `json:"email,omitempty"`
	Birthday  *time.Time `json:"birthday,omitempty"`
	LeftAt    *time.Time `json:"leftAt,omitempty"`
	JoinedAt  *time.Time `json:"joinedAt,omitempty"`
}

type UserConnection struct {
	Edges      []*db.User `json:"edges,omitempty"`
	PageInfo   *PageInfo  `json:"pageInfo"`
	TotalCount int        `json:"totalCount"`
}

type UserFilterInput struct {
	Role []*db.UserRole `json:"role,omitempty"`
}
