// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"example/pkg/db"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type AddEventCompetenceInput struct {
	EventID      string `json:"eventId"`
	CompetenceID string `json:"competenceId"`
}

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

type CompetenceSort struct {
	Field CompetenceSortField `json:"field"`
	Order SortDirection       `json:"order"`
}

type CopyFileInput struct {
	ID       string `json:"id"`
	TargetID string `json:"targetId"`
}

type CopyFilesInput struct {
	Ids      []string `json:"ids"`
	TargetID string   `json:"targetId"`
}

type CopyFilesPayload struct {
	Files []*db.File `json:"files"`
}

type CreateDomainInput struct {
	Name string `json:"name"`
}

type CreateEmailAccountInput struct {
	Name        string              `json:"name"`
	Description *string             `json:"description,omitempty"`
	Type        db.EmailAccountType `json:"type"`
	Quota       *int                `json:"quota,omitempty"`
}

type CreateEmailForwardingInput struct {
	Origin string `json:"origin"`
	Target string `json:"target"`
}

type CreateEmailGroupInput struct {
	Name        string    `json:"name"`
	Members     []*string `json:"members,omitempty"`
	Description *string   `json:"description,omitempty"`
}

type CreateEmailGroupMemberInput struct {
	Name     string `json:"name"`
	MemberOf string `json:"memberOf"`
}

type CreateEmailInput struct {
	Name    string       `json:"name"`
	Address string       `json:"address"`
	Type    db.EmailType `json:"type"`
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
	StartsAt   *string         `json:"startsAt,omitempty"`
	EndsAt     *string         `json:"endsAt,omitempty"`
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

type CreateShareInput struct {
	FileID     *string        `json:"fileId,omitempty"`
	BucketID   *string        `json:"bucketId,omitempty"`
	User       string         `json:"user"`
	Permission FilePermission `json:"permission"`
}

type CreateStudentInput struct {
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Grade     int        `json:"grade"`
	Birthday  *time.Time `json:"birthday,omitempty"`
	LeftAt    *time.Time `json:"leftAt,omitempty"`
	JoinedAt  *time.Time `json:"joinedAt,omitempty"`
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

type DeleteDomainInput struct {
	ID string `json:"id"`
}

type DeleteEmailAccountInput struct {
	ID string `json:"id"`
}

type DeleteEmailForwardingInput struct {
	ID string `json:"id"`
}

type DeleteEmailGroupMemberInput struct {
	ID string `json:"id"`
}

type DeleteEmailInput struct {
	ID string `json:"id"`
}

type DeleteFileInput struct {
	ID string `json:"id"`
}

type DeleteFilePayload struct {
	Success bool     `json:"success"`
	File    *db.File `json:"file"`
}

type DeleteFilesInput struct {
	Ids []string `json:"ids"`
}

type DeleteFilesPayload struct {
	Success bool       `json:"success"`
	Files   []*db.File `json:"files"`
}

type DeleteShareInput struct {
	FileID   *string `json:"fileId,omitempty"`
	BucketID *string `json:"bucketId,omitempty"`
	User     string  `json:"user"`
}

type DomainConnection struct {
	Edges      []*db.Domain `json:"edges,omitempty"`
	PageInfo   *PageInfo    `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

type DownloadFileInput struct {
	ID string `json:"id"`
}

type DownloadFilePayload struct {
	URL string `json:"url"`
}

type DownloadFilesInput struct {
	Ids []string `json:"ids"`
}

type DownloadFilesPayload struct {
	// The url to download a zip file containing all the files.
	URL string `json:"url"`
}

type EmailAccountConnection struct {
	Edges      []*db.EmailAccount `json:"edges,omitempty"`
	PageInfo   *PageInfo          `json:"pageInfo"`
	TotalCount int                `json:"totalCount"`
}

type EmailAccountFilter struct {
	Type *db.EmailAccountType `json:"type,omitempty"`
}

type EmailConnection struct {
	Edges      []*db.Email `json:"edges,omitempty"`
	PageInfo   *PageInfo   `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

type EmailForwardingConnection struct {
	Edges      []*db.EmailForwarding `json:"edges,omitempty"`
	PageInfo   *PageInfo             `json:"pageInfo"`
	TotalCount int                   `json:"totalCount"`
}

type EmailGroupMemberConnection struct {
	Edges      []*db.EmailGroupMember `json:"edges,omitempty"`
	PageInfo   *PageInfo              `json:"pageInfo"`
	TotalCount int                    `json:"totalCount"`
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
	MyBucket *bool   `json:"myBucket,omitempty"`
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

type MoveFileInput struct {
	ID       string `json:"id"`
	TargetID string `json:"targetId"`
}

type MoveFilesInput struct {
	Ids      []string `json:"ids"`
	TargetID string   `json:"targetId"`
}

type MoveFilesPayload struct {
	Files []*db.File `json:"files"`
}

type MyFilesFilterInput struct {
	ParentID *string `json:"parentId,omitempty"`
}

type PageInfo struct {
	HasNextPage     bool `json:"hasNextPage"`
	HasPreviousPage bool `json:"hasPreviousPage"`
	CurrentPage     int  `json:"currentPage"`
}

type PreviewFileInput struct {
	ID string `json:"id"`
}

type PreviewFilePayload struct {
	URL string `json:"url"`
}

type RenameFileInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RenameSharedDriveInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ReportConnection struct {
	Edges      []*db.Report `json:"edges,omitempty"`
	PageInfo   *PageInfo    `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

type ResetPasswordInput struct {
	Token    *string `json:"token,omitempty"`
	Password string  `json:"password"`
}

type ResetPasswordPayload struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ShareFileInput struct {
	FileID     string         `json:"fileId"`
	Users      []string       `json:"users"`
	Emails     []string       `json:"emails"`
	Permission FilePermission `json:"permission"`
}

type ShareInput struct {
	BucketID *string `json:"bucketId,omitempty"`
	FileID   *string `json:"fileId,omitempty"`
}

type ShareUser struct {
	User       *db.User       `json:"user"`
	Permission FilePermission `json:"permission"`
}

type SharedDriveFilterInput struct {
	Folder *string `json:"folder,omitempty"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInPayload struct {
	Token         string   `json:"token"`
	EnabledApps   []string `json:"enabled_apps"`
	Language      string   `json:"language"`
	SetupComplete bool     `json:"setupComplete"`
}

type SignUpInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type SortCompetenceInput struct {
	ID        string `json:"id"`
	SortOrder int    `json:"sortOrder"`
}

type TagConnection struct {
	Edges      []*db.Tag `json:"edges,omitempty"`
	PageInfo   *PageInfo `json:"pageInfo"`
	TotalCount int       `json:"totalCount"`
}

type UpdateCompetenceInput struct {
	ID    string  `json:"id"`
	Color *string `json:"color,omitempty"`
}

type UpdateCompetenceSortingInput struct {
	Competences []*SortCompetenceInput `json:"competences"`
}

type UpdateEmailAccountInput struct {
	ID          string               `json:"id"`
	Name        *string              `json:"name,omitempty"`
	Description *string              `json:"description,omitempty"`
	Type        *db.EmailAccountType `json:"type,omitempty"`
	Quota       *int                 `json:"quota,omitempty"`
	Active      *bool                `json:"active,omitempty"`
}

type UpdateEmailGroupInput struct {
	ID          string    `json:"id"`
	Name        *string   `json:"name,omitempty"`
	Members     []*string `json:"members,omitempty"`
	Description *string   `json:"description,omitempty"`
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
	StartsAt   *string         `json:"startsAt,omitempty"`
	EndsAt     *string         `json:"endsAt,omitempty"`
	Recurrence []*string       `json:"recurrence,omitempty"`
}

type UpdateOrganisationInput struct {
	ID        string  `json:"id"`
	Name      *string `json:"name,omitempty"`
	LegalName *string `json:"legalName,omitempty"`
	Website   *string `json:"website,omitempty"`
	Phone     *string `json:"phone,omitempty"`
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
	Grade     *int       `json:"grade,omitempty"`
	Birthday  *time.Time `json:"birthday,omitempty"`
	LeftAt    *time.Time `json:"leftAt,omitempty"`
	JoinedAt  *time.Time `json:"joinedAt,omitempty"`
}

type UploadFilesPayload struct {
	Files []*db.File `json:"files"`
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
	Role        []*db.UserRole `json:"role,omitempty"`
	OrderBy     *UserOrderBy   `json:"orderBy,omitempty"`
	ShowDeleted *bool          `json:"showDeleted,omitempty"`
}

type UserStudentConnection struct {
	Edges      []*db.UserStudent `json:"edges,omitempty"`
	PageInfo   *PageInfo         `json:"pageInfo"`
	TotalCount int               `json:"totalCount"`
}

type CompetenceSortField string

const (
	CompetenceSortFieldCreatedAt CompetenceSortField = "created_at"
	CompetenceSortFieldName      CompetenceSortField = "name"
	CompetenceSortFieldSortOrder CompetenceSortField = "sort_order"
)

var AllCompetenceSortField = []CompetenceSortField{
	CompetenceSortFieldCreatedAt,
	CompetenceSortFieldName,
	CompetenceSortFieldSortOrder,
}

func (e CompetenceSortField) IsValid() bool {
	switch e {
	case CompetenceSortFieldCreatedAt, CompetenceSortFieldName, CompetenceSortFieldSortOrder:
		return true
	}
	return false
}

func (e CompetenceSortField) String() string {
	return string(e)
}

func (e *CompetenceSortField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CompetenceSortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CompetenceSortField", str)
	}
	return nil
}

func (e CompetenceSortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EntrySortBy string

const (
	EntrySortByDateAsc       EntrySortBy = "date_ASC"
	EntrySortByDateDesc      EntrySortBy = "date_DESC"
	EntrySortByCreatedAtAsc  EntrySortBy = "createdAt_ASC"
	EntrySortByCreatedAtDesc EntrySortBy = "createdAt_DESC"
)

var AllEntrySortBy = []EntrySortBy{
	EntrySortByDateAsc,
	EntrySortByDateDesc,
	EntrySortByCreatedAtAsc,
	EntrySortByCreatedAtDesc,
}

func (e EntrySortBy) IsValid() bool {
	switch e {
	case EntrySortByDateAsc, EntrySortByDateDesc, EntrySortByCreatedAtAsc, EntrySortByCreatedAtDesc:
		return true
	}
	return false
}

func (e EntrySortBy) String() string {
	return string(e)
}

func (e *EntrySortBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EntrySortBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EntrySortBy", str)
	}
	return nil
}

func (e EntrySortBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EventOrderBy string

const (
	EventOrderByStartsAtAsc  EventOrderBy = "startsAt_ASC"
	EventOrderByStartsAtDesc EventOrderBy = "startsAt_DESC"
	EventOrderByEndsAtAsc    EventOrderBy = "endsAt_ASC"
	EventOrderByEndsAtDesc   EventOrderBy = "endsAt_DESC"
)

var AllEventOrderBy = []EventOrderBy{
	EventOrderByStartsAtAsc,
	EventOrderByStartsAtDesc,
	EventOrderByEndsAtAsc,
	EventOrderByEndsAtDesc,
}

func (e EventOrderBy) IsValid() bool {
	switch e {
	case EventOrderByStartsAtAsc, EventOrderByStartsAtDesc, EventOrderByEndsAtAsc, EventOrderByEndsAtDesc:
		return true
	}
	return false
}

func (e EventOrderBy) String() string {
	return string(e)
}

func (e *EventOrderBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EventOrderBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EventOrderBy", str)
	}
	return nil
}

func (e EventOrderBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FilePermission string

const (
	FilePermissionViewer  FilePermission = "Viewer"
	FilePermissionManager FilePermission = "Manager"
)

var AllFilePermission = []FilePermission{
	FilePermissionViewer,
	FilePermissionManager,
}

func (e FilePermission) IsValid() bool {
	switch e {
	case FilePermissionViewer, FilePermissionManager:
		return true
	}
	return false
}

func (e FilePermission) String() string {
	return string(e)
}

func (e *FilePermission) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FilePermission(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FilePermission", str)
	}
	return nil
}

func (e FilePermission) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SortDirection string

const (
	SortDirectionAsc  SortDirection = "asc"
	SortDirectionDesc SortDirection = "desc"
)

var AllSortDirection = []SortDirection{
	SortDirectionAsc,
	SortDirectionDesc,
}

func (e SortDirection) IsValid() bool {
	switch e {
	case SortDirectionAsc, SortDirectionDesc:
		return true
	}
	return false
}

func (e SortDirection) String() string {
	return string(e)
}

func (e *SortDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortDirection", str)
	}
	return nil
}

func (e SortDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserOrderBy string

const (
	UserOrderByFirstNameAsc  UserOrderBy = "firstNameAsc"
	UserOrderByFirstNameDesc UserOrderBy = "firstNameDesc"
	UserOrderByLastNameAsc   UserOrderBy = "lastNameAsc"
	UserOrderByLastNameDesc  UserOrderBy = "lastNameDesc"
)

var AllUserOrderBy = []UserOrderBy{
	UserOrderByFirstNameAsc,
	UserOrderByFirstNameDesc,
	UserOrderByLastNameAsc,
	UserOrderByLastNameDesc,
}

func (e UserOrderBy) IsValid() bool {
	switch e {
	case UserOrderByFirstNameAsc, UserOrderByFirstNameDesc, UserOrderByLastNameAsc, UserOrderByLastNameDesc:
		return true
	}
	return false
}

func (e UserOrderBy) String() string {
	return string(e)
}

func (e *UserOrderBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserOrderBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserOrderBy", str)
	}
	return nil
}

func (e UserOrderBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
