// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: entry.sql

package db

import (
	"context"
)

const getEntry = `-- name: GetEntry :one
SELECT id, date, body, user_id, created_at, deleted_at, organisation_id
FROM entries
WHERE id = $1 AND organisation_id = $2
`

type GetEntryParams struct {
	ID             string `json:"id"`
	OrganisationID string `json:"organisation_id"`
}

func (q *Queries) GetEntry(ctx context.Context, arg GetEntryParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntry, arg.ID, arg.OrganisationID)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.Date,
		&i.Body,
		&i.UserID,
		&i.CreatedAt,
		&i.DeletedAt,
		&i.OrganisationID,
	)
	return i, err
}

const listEntries = `-- name: ListEntries :many
SELECT id, date, body, user_id, created_at, deleted_at, organisation_id
FROM entries
WHERE entries.organisation_id = $1
LIMIT $2 OFFSET $3
`

type ListEntriesParams struct {
	OrganisationID string `json:"organisation_id"`
	Limit          int32  `json:"limit"`
	Offset         int32  `json:"offset"`
}

func (q *Queries) ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, listEntries, arg.OrganisationID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Entry
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.Date,
			&i.Body,
			&i.UserID,
			&i.CreatedAt,
			&i.DeletedAt,
			&i.OrganisationID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
