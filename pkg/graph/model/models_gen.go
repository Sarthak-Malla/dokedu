// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"example/pkg/db"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type BucketConnection struct {
	Edges      []*db.Bucket `json:"edges"`
	TotalCount int          `json:"totalCount"`
	PageInfo   *PageInfo    `json:"pageInfo"`
}

type BucketFilterInput struct {
	Shared *bool `json:"shared,omitempty"`
}

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
	UserID  *string              `json:"userId,omitempty"`
}

type CreateEntryInput struct {
	Date            string                       `json:"date"`
	Body            string                       `json:"body"`
	TagIds          []string                     `json:"tagIds,omitempty"`
	FileIds         []string                     `json:"fileIds,omitempty"`
	UserIds         []string                     `json:"userIds,omitempty"`
	EventIds        []string                     `json:"eventIds,omitempty"`
	UserCompetences []*CreateUserCompetenceInput `json:"userCompetences,omitempty"`
}

type CreateEventInput struct {
	Title      string          `json:"title"`
	Image      *graphql.Upload `json:"image,omitempty"`
	Body       *string         `json:"body,omitempty"`
	StartsAt   time.Time       `json:"startsAt"`
	EndsAt     time.Time       `json:"endsAt"`
	Recurrence []*string       `json:"recurrence,omitempty"`
}

type CreateFolderInput struct {
	Name     string  `json:"name"`
	ParentID *string `json:"parentId,omitempty"`
	BucketID *string `json:"bucketId,omitempty"`
}

type CreateReportInput struct {
	Format      db.ReportFormat `json:"format"`
	Kind        db.ReportKind   `json:"kind"`
	From        time.Time       `json:"from"`
	To          time.Time       `json:"to"`
	FilterTags  []string        `json:"filterTags"`
	StudentUser string          `json:"studentUser"`
}

type CreateTagInput struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type CreateUserCompetenceInput struct {
	Level        int    `json:"level"`
	UserID       string `json:"userId"`
	CompetenceID string `json:"competenceId"`
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
	Authors     []*string  `json:"authors,omitempty"`
	Users       []*string  `json:"users,omitempty"`
	Tags        []*string  `json:"tags,omitempty"`
	Competences []*string  `json:"competences,omitempty"`
	From        *time.Time `json:"from,omitempty"`
	To          *time.Time `json:"to,omitempty"`
	Deleted     *bool      `json:"deleted,omitempty"`
}

type EventConnection struct {
	Edges      []*db.Event `json:"edges,omitempty"`
	PageInfo   *PageInfo   `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

type EventFilterInput struct {
	From    *time.Time `json:"from,omitempty"`
	To      *time.Time `json:"to,omitempty"`
	Deleted *bool      `json:"deleted,omitempty"`
}

type ExportEventsInput struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Deleted bool   `json:"deleted"`
}

type ExportEventsPayload struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	StartsAt string `json:"startsAt"`
	EndsAt   string `json:"endsAt"`
	Subjects string `json:"subjects"`
}

type FileConnection struct {
	Edges      []*db.File `json:"edges"`
	TotalCount int        `json:"totalCount"`
	PageInfo   *PageInfo  `json:"pageInfo"`
}

type FileUploadInput struct {
	File graphql.Upload `json:"file"`
	// The folder to upload the file to if empty the file will be uploaded to the root folder of the user.
	ParentID *string `json:"parentId,omitempty"`
	// The shared drive to upload the file to if empty the file will be uploaded to the root folder of the user.
	BucketID *string `json:"bucketId,omitempty"`
}

type FilesFilterInput struct {
	ParentID *string `json:"parentId,omitempty"`
	BucketID *string `json:"bucketId,omitempty"`
	Limit    *int    `json:"limit,omitempty"`
	Offset   *int    `json:"offset,omitempty"`
}

type ForgotPasswordInput struct {
	Email string `json:"email"`
}

type ForgotPasswordPayload struct {
	Success bool `json:"success"`
}

type GenerateFileURLInput struct {
	ID string `json:"id"`
}

type GenerateFileURLPayload struct {
	URL string `json:"url"`
}

type MyFilesFilterInput struct {
	ParentID *string `json:"parentId,omitempty"`
	Limit    *int    `json:"limit,omitempty"`
	Offset   *int    `json:"offset,omitempty"`
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

type ResetPasswordInput struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

type ResetPasswordPayload struct {
	Success bool `json:"success"`
}

type SharedDriveFilterInput struct {
	Folder *string `json:"folder,omitempty"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInPayload struct {
	Token       string   `json:"token"`
	EnabledApps []string `json:"enabled_apps"`
}

type SignUpInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UpdateCompetenceInput struct {
	ID    string  `json:"id"`
	Color *string `json:"color,omitempty"`
}

type UpdateEntryInput struct {
	ID              string                       `json:"id"`
	Date            *string                      `json:"date,omitempty"`
	Body            *string                      `json:"body,omitempty"`
	TagIds          []string                     `json:"tagIds,omitempty"`
	FileIds         []string                     `json:"fileIds,omitempty"`
	UserIds         []string                     `json:"userIds,omitempty"`
	EventIds        []string                     `json:"eventIds,omitempty"`
	UserCompetences []*UpdateUserCompetenceInput `json:"userCompetences,omitempty"`
}

type UpdateEventInput struct {
	ID         string          `json:"id"`
	Title      *string         `json:"title,omitempty"`
	Image      *graphql.Upload `json:"image,omitempty"`
	Body       *string         `json:"body,omitempty"`
	StartsAt   *time.Time      `json:"startsAt,omitempty"`
	EndsAt     *time.Time      `json:"endsAt,omitempty"`
	Recurrence []*string       `json:"recurrence,omitempty"`
}

type UpdateUserCompetenceInput struct {
	Level        int    `json:"level"`
	UserID       string `json:"userId"`
	CompetenceID string `json:"competenceId"`
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

type UserCompetenceConnection struct {
	Edges      []*db.UserCompetence `json:"edges,omitempty"`
	PageInfo   *PageInfo            `json:"pageInfo"`
	TotalCount int                  `json:"totalCount"`
}

type UserCompetenceFilterInput struct {
	UserID       *string `json:"userID,omitempty"`
	CompetenceID *string `json:"competenceID,omitempty"`
}

type UserConnection struct {
	Edges      []*db.User `json:"edges,omitempty"`
	PageInfo   *PageInfo  `json:"pageInfo"`
	TotalCount int        `json:"totalCount"`
}

type UserFileFilterInput struct {
	FolderID *string `json:"folderId,omitempty"`
}

type UserFilterInput struct {
	Role []*db.UserRole `json:"role,omitempty"`
}

type UserStudentConnection struct {
	Edges      []*db.UserStudent `json:"edges,omitempty"`
	PageInfo   *PageInfo         `json:"pageInfo"`
	TotalCount int               `json:"totalCount"`
}
