package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/dokedu/dokedu/backend/pkg/graph/model"
	"github.com/dokedu/dokedu/backend/pkg/middleware"
	"github.com/dokedu/dokedu/backend/pkg/msg"
	"github.com/dokedu/dokedu/backend/pkg/services/database"
	"github.com/dokedu/dokedu/backend/pkg/services/database/db"
)

// Type is the resolver for the type field.
func (r *competenceResolver) Type(ctx context.Context, obj *db.Competence) (db.CompetenceType, error) {
	panic(fmt.Errorf("not implemented: Type - type"))
}

// Color is the resolver for the color field.
func (r *competenceResolver) Color(ctx context.Context, obj *db.Competence) (string, error) {
	panic(fmt.Errorf("not implemented: Color - color"))
}

// Parents is the resolver for the parents field.
func (r *competenceResolver) Parents(ctx context.Context, obj *db.Competence) ([]db.Competence, error) {
	panic(fmt.Errorf("not implemented: Parents - parents"))
}

// SortOrder is the resolver for the sortOrder field.
func (r *competenceResolver) SortOrder(ctx context.Context, obj *db.Competence) (int, error) {
	panic(fmt.Errorf("not implemented: SortOrder - sortOrder"))
}

// Competences is the resolver for the competences field.
func (r *competenceResolver) Competences(ctx context.Context, obj *db.Competence, search *string, sort *model.CompetenceSort) ([]*db.Competence, error) {
	panic(fmt.Errorf("not implemented: Competences - competences"))
}

// UserCompetences is the resolver for the userCompetences field.
func (r *competenceResolver) UserCompetences(ctx context.Context, obj *db.Competence, userID *string) ([]*db.UserCompetence, error) {
	panic(fmt.Errorf("not implemented: UserCompetences - userCompetences"))
}

// Tendency is the resolver for the tendency field.
func (r *competenceResolver) Tendency(ctx context.Context, obj *db.Competence, userID string) (*model.CompetenceTendency, error) {
	panic(fmt.Errorf("not implemented: Tendency - tendency"))
}

// SignIn is the resolver for the signIn field.
func (r *mutationResolver) SignIn(ctx context.Context, input model.SignInInput) (*model.SignInPayload, error) {
	user, err := r.DB.GLOBAL_UserFindByEmail(ctx, input.Email)
	if err != nil {
		return nil, msg.ErrInvalidEmailOrPassword
	}

	// user.Password is a sql.NullString, so we need to check if it is valid
	if !user.Password.Valid {
		return nil, msg.ErrInvalidEmailOrPassword
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(input.Password)); err != nil {
		return nil, msg.ErrInvalidEmailOrPassword
	}

	// Generate a new token
	token := gonanoid.Must(32)
	_, err = r.DB.GLOBAL_SessionCreate(ctx, db.GLOBAL_SessionCreateParams{
		UserID: user.ID,
		Token:  token,
	})
	if err != nil {
		return nil, errors.New("unable to generate a token")
	}

	// todo: return more data
	return &model.SignInPayload{
		Token: token,
		User:  user,
	}, nil
}

// ResetPassword is the resolver for the resetPassword field.
func (r *mutationResolver) ResetPassword(ctx context.Context, input model.ResetPasswordInput) (*model.ResetPasswordPayload, error) {
	_, ok := middleware.GetUser(ctx)
	if ok {
		return nil, msg.ErrInvalidInput
	}

	// if there is no user, try to find by recover token
	if input.Token == nil {
		return nil, msg.ErrInvalidInput
	}

	var err error
	user, err := r.DB.GLOBAL_UserFindByRecoveryToken(ctx, *input.Token)
	if err != nil {
		return nil, msg.ErrInvalidRecoveryToken
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	_, err = r.DB.UserUpdatePassword(ctx, db.UserUpdatePasswordParams{
		Password:       string(hashedPassword),
		ID:             user.ID,
		OrganisationID: user.OrganisationID,
	})
	if err != nil {
		return nil, errors.New("unable to update the password")
	}

	// todo: resetting your password should probably invalidate other sessions

	return &model.ResetPasswordPayload{
		Success: true,
	}, nil
}

// ForgotPassword is the resolver for the forgotPassword field.
func (r *mutationResolver) ForgotPassword(ctx context.Context, input model.ForgotPasswordInput) (*model.ForgotPasswordPayload, error) {
	user, err := r.DB.GLOBAL_UserFindByEmail(ctx, input.Email)
	if err != nil {
		return nil, msg.ErrInvalidInput
	}

	// Generate a new token
	token := gonanoid.Must(32)

	// Update the user with the new token
	_, err = r.DB.UserUpdateRecoveryToken(ctx, db.UserUpdateRecoveryTokenParams{
		RecoveryToken:  token,
		ID:             user.ID,
		OrganisationID: user.OrganisationID,
	})
	if err != nil {
		return nil, err
	}

	// Send the email to the user
	err = r.Services.Mail.SendPasswordReset(input.Email, user.FirstName, user.LanguageOrDefault(), token)
	if err != nil {
		return nil, err
	}

	return &model.ForgotPasswordPayload{Success: true}, nil
}

// SignOut is the resolver for the signOut field.
func (r *mutationResolver) SignOut(ctx context.Context) (bool, error) {
	user, ok := middleware.GetUser(ctx)
	if !ok {
		return false, msg.ErrUnauthorized
	}

	err := r.DB.GLOBAL_SessionsDeleteByUserID(ctx, user.User.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}

// AcceptInvite is the resolver for the acceptInvite field.
func (r *mutationResolver) AcceptInvite(ctx context.Context, token string, input model.SignUpInput) (*model.SignInPayload, error) {
	panic(fmt.Errorf("not implemented: AcceptInvite - acceptInvite"))
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*db.User, error) {
	user, ok := middleware.GetUser(ctx)
	if !ok || !user.HasPermissionAdmin() {
		return nil, msg.ErrUnauthorized
	}

	createdUser, err := r.DB.UserCreate(ctx, db.UserCreateParams{
		Role:           input.Role,
		OrganisationID: user.OrganisationID,
		FirstName:      input.FirstName,
		LastName:       input.LastName,
		Email:          pgtype.Text{Valid: true, String: input.Email},
	})
	if err != nil {
		return nil, err
	}

	token := gonanoid.Must(32)
	_, err = r.DB.UserUpdateRecoveryToken(ctx, db.UserUpdateRecoveryTokenParams{
		RecoveryToken:  token,
		ID:             createdUser.ID,
		OrganisationID: user.OrganisationID,
	})
	if err != nil {
		return nil, err
	}

	err = r.Services.Mail.SendInvite(input.Email, createdUser.FirstName, user.FirstName, user.LanguageOrDefault(), token)
	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*db.User, error) {
	user, ok := middleware.GetUser(ctx)
	if !ok || !user.HasPermissionAdmin() {
		return nil, msg.ErrUnauthorized
	}

	// update the common user fields
	updatedUser, err := r.DB.UserUpdate(ctx, db.UserUpdateParams{
		FirstName:      input.FirstName,
		LastName:       input.LastName,
		ID:             input.ID,
		OrganisationID: user.OrganisationID,
	})
	if err != nil {
		return nil, err
	}

	// if the updated user is a student, update the student fields as well
	if updatedUser.Role == db.UserRoleStudent {
		query := r.DB.NewQueryBuilder().
			Update("user_students").
			Where("user_id = ? and organisation_id = ?", updatedUser.ID, user.OrganisationID)

		var shouldUpdate bool
		if input.Grade != nil {
			shouldUpdate = true
			query = query.Set("grade", *input.Grade)
		}
		if input.Birthday != nil {
			shouldUpdate = true
			query = query.Set("birthday", *input.Birthday)
		}
		if input.LeftAt != nil {
			shouldUpdate = true
			query = query.Set("left_at", *input.LeftAt)
		}
		if input.JoinedAt != nil {
			shouldUpdate = true
			query = query.Set("joined_at", *input.JoinedAt)
		}
		if input.Emoji != nil {
			shouldUpdate = true
			query = query.Set("emoji", *input.Emoji)
		}
		if input.MissedHours != nil {
			shouldUpdate = true
			query = query.Set("missed_hours", *input.MissedHours)
		}
		if input.MissedHoursExcused != nil {
			shouldUpdate = true
			query = query.Set("missed_hours_excused", *input.MissedHoursExcused)
		}

		if shouldUpdate {
			err = database.ExecUpdate(r.DB, ctx, query)
			if err != nil {
				return nil, err
			}
		}
	}

	return &updatedUser, nil
}

// ArchiveUser is the resolver for the archiveUser field.
func (r *mutationResolver) ArchiveUser(ctx context.Context, id string) (*db.User, error) {
	user, ok := middleware.GetUser(ctx)
	if !ok || !user.HasPermissionAdmin() {
		return nil, msg.ErrUnauthorized
	}

	// try to find the user
	userToBeArchived, err := r.DB.UserFindByID(ctx, db.UserFindByIDParams{
		ID:             id,
		OrganisationID: user.OrganisationID,
	})
	if err != nil {
		return nil, err
	}

	// archive the user
	_, err = r.DB.UserSoftDelete(ctx, db.UserSoftDeleteParams{
		ID:             id,
		OrganisationID: user.OrganisationID,
	})
	if err != nil {
		return nil, err
	}

	// if the user is a student, we need to delete the student relation as well
	if userToBeArchived.Role == db.UserRoleStudent {
		_, err = r.DB.UserStudentSoftDeleteByUserID(ctx, db.UserStudentSoftDeleteByUserIDParams{
			UserID:         id,
			OrganisationID: user.OrganisationID,
		})
		if err != nil {
			return nil, err
		}
	}

	// delete all sessions for the user
	err = r.DB.GLOBAL_SessionsDeleteByUserID(ctx, userToBeArchived.ID)
	if err != nil {
		return nil, err
	}

	return &userToBeArchived, nil
}

// UpdateUserLanguage is the resolver for the updateUserLanguage field.
func (r *mutationResolver) UpdateUserLanguage(ctx context.Context, language model.UserLanguage) (*db.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUserLanguage - updateUserLanguage"))
}

// SendUserInvite is the resolver for the sendUserInvite field.
func (r *mutationResolver) SendUserInvite(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: SendUserInvite - sendUserInvite"))
}

// CreateStudent is the resolver for the createStudent field.
func (r *mutationResolver) CreateStudent(ctx context.Context, input model.CreateStudentInput) (*db.User, error) {
	panic(fmt.Errorf("not implemented: CreateStudent - createStudent"))
}

// CreateUserCompetence is the resolver for the createUserCompetence field.
func (r *mutationResolver) CreateUserCompetence(ctx context.Context, input model.CreateUserCompetenceInput) (*db.UserCompetence, error) {
	panic(fmt.Errorf("not implemented: CreateUserCompetence - createUserCompetence"))
}

// ArchiveUserCompetence is the resolver for the archiveUserCompetence field.
func (r *mutationResolver) ArchiveUserCompetence(ctx context.Context, id string) (*db.UserCompetence, error) {
	panic(fmt.Errorf("not implemented: ArchiveUserCompetence - archiveUserCompetence"))
}

// CreateTag is the resolver for the createTag field.
func (r *mutationResolver) CreateTag(ctx context.Context, input model.CreateTagInput) (*db.Tag, error) {
	panic(fmt.Errorf("not implemented: CreateTag - createTag"))
}

// ArchiveTag is the resolver for the archiveTag field.
func (r *mutationResolver) ArchiveTag(ctx context.Context, id string) (*db.Tag, error) {
	panic(fmt.Errorf("not implemented: ArchiveTag - archiveTag"))
}

// UpdateTag is the resolver for the updateTag field.
func (r *mutationResolver) UpdateTag(ctx context.Context, id string, input model.CreateTagInput) (*db.Tag, error) {
	panic(fmt.Errorf("not implemented: UpdateTag - updateTag"))
}

// UpdatePassword is the resolver for the updatePassword field.
func (r *mutationResolver) UpdatePassword(ctx context.Context, oldPassword string, newPassword string) (bool, error) {
	panic(fmt.Errorf("not implemented: UpdatePassword - updatePassword"))
}

// UpdateCompetence is the resolver for the updateCompetence field.
func (r *mutationResolver) UpdateCompetence(ctx context.Context, input model.UpdateCompetenceInput) (*db.Competence, error) {
	panic(fmt.Errorf("not implemented: UpdateCompetence - updateCompetence"))
}

// UpdateCompetenceSorting is the resolver for the updateCompetenceSorting field.
func (r *mutationResolver) UpdateCompetenceSorting(ctx context.Context, input model.UpdateCompetenceSortingInput) ([]*db.Competence, error) {
	panic(fmt.Errorf("not implemented: UpdateCompetenceSorting - updateCompetenceSorting"))
}

// CreateCompetence is the resolver for the createCompetence field.
func (r *mutationResolver) CreateCompetence(ctx context.Context, input model.CreateCompetenceInput) (*db.Competence, error) {
	panic(fmt.Errorf("not implemented: CreateCompetence - createCompetence"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int, filter *model.UserFilterInput, search *string) (*model.UserConnection, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*db.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*db.User, error) {
	panic(fmt.Errorf("not implemented: Me - me"))
}

// Competence is the resolver for the competence field.
func (r *queryResolver) Competence(ctx context.Context, id string) (*db.Competence, error) {
	panic(fmt.Errorf("not implemented: Competence - competence"))
}

// Competences is the resolver for the competences field.
func (r *queryResolver) Competences(ctx context.Context, limit *int, offset *int, filter *model.CompetenceFilterInput, search *string, sort *model.CompetenceSort) (*model.CompetenceConnection, error) {
	panic(fmt.Errorf("not implemented: Competences - competences"))
}

// Tag is the resolver for the tag field.
func (r *queryResolver) Tag(ctx context.Context, id string) (*db.Tag, error) {
	panic(fmt.Errorf("not implemented: Tag - tag"))
}

// Tags is the resolver for the tags field.
func (r *queryResolver) Tags(ctx context.Context, limit *int, offset *int, search *string) (*model.TagConnection, error) {
	panic(fmt.Errorf("not implemented: Tags - tags"))
}

// UserStudents is the resolver for the userStudents field.
func (r *queryResolver) UserStudents(ctx context.Context, limit *int, offset *int) (*model.UserStudentConnection, error) {
	panic(fmt.Errorf("not implemented: UserStudents - userStudents"))
}

// UserStudent is the resolver for the userStudent field.
func (r *queryResolver) UserStudent(ctx context.Context, id string) (*db.UserStudent, error) {
	panic(fmt.Errorf("not implemented: UserStudent - userStudent"))
}

// InviteDetails is the resolver for the inviteDetails field.
func (r *queryResolver) InviteDetails(ctx context.Context, token string) (*model.InviteDetailsPayload, error) {
	panic(fmt.Errorf("not implemented: InviteDetails - inviteDetails"))
}

// Color is the resolver for the color field.
func (r *tagResolver) Color(ctx context.Context, obj *db.Tag) (string, error) {
	panic(fmt.Errorf("not implemented: Color - color"))
}

// DeletedAt is the resolver for the deletedAt field.
func (r *tagResolver) DeletedAt(ctx context.Context, obj *db.Tag) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: DeletedAt - deletedAt"))
}

// Email is the resolver for the email field.
func (r *userResolver) Email(ctx context.Context, obj *db.User) (*string, error) {
	panic(fmt.Errorf("not implemented: Email - email"))
}

// Student is the resolver for the student field.
func (r *userResolver) Student(ctx context.Context, obj *db.User) (*db.UserStudent, error) {
	panic(fmt.Errorf("not implemented: Student - student"))
}

// Language is the resolver for the language field.
func (r *userResolver) Language(ctx context.Context, obj *db.User) (*model.UserLanguage, error) {
	panic(fmt.Errorf("not implemented: Language - language"))
}

// DeletedAt is the resolver for the deletedAt field.
func (r *userResolver) DeletedAt(ctx context.Context, obj *db.User) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: DeletedAt - deletedAt"))
}

// InviteAccepted is the resolver for the inviteAccepted field.
func (r *userResolver) InviteAccepted(ctx context.Context, obj *db.User) (bool, error) {
	panic(fmt.Errorf("not implemented: InviteAccepted - inviteAccepted"))
}

// LastSeenAt is the resolver for the lastSeenAt field.
func (r *userResolver) LastSeenAt(ctx context.Context, obj *db.User) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: LastSeenAt - lastSeenAt"))
}

// Competence is the resolver for the competence field.
func (r *userCompetenceResolver) Competence(ctx context.Context, obj *db.UserCompetence) (*db.Competence, error) {
	panic(fmt.Errorf("not implemented: Competence - competence"))
}

// Entry is the resolver for the entry field.
func (r *userCompetenceResolver) Entry(ctx context.Context, obj *db.UserCompetence) (*db.Entry, error) {
	panic(fmt.Errorf("not implemented: Entry - entry"))
}

// User is the resolver for the user field.
func (r *userCompetenceResolver) User(ctx context.Context, obj *db.UserCompetence) (*db.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// CreatedBy is the resolver for the createdBy field.
func (r *userCompetenceResolver) CreatedBy(ctx context.Context, obj *db.UserCompetence) (*db.User, error) {
	panic(fmt.Errorf("not implemented: CreatedBy - createdBy"))
}

// LeftAt is the resolver for the leftAt field.
func (r *userStudentResolver) LeftAt(ctx context.Context, obj *db.UserStudent) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: LeftAt - leftAt"))
}

// Birthday is the resolver for the birthday field.
func (r *userStudentResolver) Birthday(ctx context.Context, obj *db.UserStudent) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: Birthday - birthday"))
}

// Nationality is the resolver for the nationality field.
func (r *userStudentResolver) Nationality(ctx context.Context, obj *db.UserStudent) (*string, error) {
	panic(fmt.Errorf("not implemented: Nationality - nationality"))
}

// Comments is the resolver for the comments field.
func (r *userStudentResolver) Comments(ctx context.Context, obj *db.UserStudent) (*string, error) {
	panic(fmt.Errorf("not implemented: Comments - comments"))
}

// JoinedAt is the resolver for the joinedAt field.
func (r *userStudentResolver) JoinedAt(ctx context.Context, obj *db.UserStudent) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: JoinedAt - joinedAt"))
}

// DeletedAt is the resolver for the deletedAt field.
func (r *userStudentResolver) DeletedAt(ctx context.Context, obj *db.UserStudent) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: DeletedAt - deletedAt"))
}

// EntriesCount is the resolver for the entriesCount field.
func (r *userStudentResolver) EntriesCount(ctx context.Context, obj *db.UserStudent) (int, error) {
	panic(fmt.Errorf("not implemented: EntriesCount - entriesCount"))
}

// CompetencesCount is the resolver for the competencesCount field.
func (r *userStudentResolver) CompetencesCount(ctx context.Context, obj *db.UserStudent) (int, error) {
	panic(fmt.Errorf("not implemented: CompetencesCount - competencesCount"))
}

// EventsCount is the resolver for the eventsCount field.
func (r *userStudentResolver) EventsCount(ctx context.Context, obj *db.UserStudent) (int, error) {
	panic(fmt.Errorf("not implemented: EventsCount - eventsCount"))
}

// Emoji is the resolver for the emoji field.
func (r *userStudentResolver) Emoji(ctx context.Context, obj *db.UserStudent) (*string, error) {
	panic(fmt.Errorf("not implemented: Emoji - emoji"))
}

// User is the resolver for the user field.
func (r *userStudentResolver) User(ctx context.Context, obj *db.UserStudent) (*db.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// MissedHours is the resolver for the missedHours field.
func (r *userStudentResolver) MissedHours(ctx context.Context, obj *db.UserStudent) (int, error) {
	panic(fmt.Errorf("not implemented: MissedHours - missedHours"))
}

// MissedHoursExcused is the resolver for the missedHoursExcused field.
func (r *userStudentResolver) MissedHoursExcused(ctx context.Context, obj *db.UserStudent) (int, error) {
	panic(fmt.Errorf("not implemented: MissedHoursExcused - missedHoursExcused"))
}

// Competence returns CompetenceResolver implementation.
func (r *Resolver) Competence() CompetenceResolver { return &competenceResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Tag returns TagResolver implementation.
func (r *Resolver) Tag() TagResolver { return &tagResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

// UserCompetence returns UserCompetenceResolver implementation.
func (r *Resolver) UserCompetence() UserCompetenceResolver { return &userCompetenceResolver{r} }

// UserStudent returns UserStudentResolver implementation.
func (r *Resolver) UserStudent() UserStudentResolver { return &userStudentResolver{r} }

type competenceResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type tagResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type userCompetenceResolver struct{ *Resolver }
type userStudentResolver struct{ *Resolver }
