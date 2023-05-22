// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: tag.sql

package db

import (
	"context"
)

const getTag = `-- name: GetTag :one
SELECT id, name, color, organisation_id, created_at, deleted_at
FROM tags
WHERE id = $1 AND organisation_id = $2
`

type GetTagParams struct {
	ID             string `json:"id"`
	OrganisationID string `json:"organisation_id"`
}

func (q *Queries) GetTag(ctx context.Context, arg GetTagParams) (Tag, error) {
	row := q.db.QueryRowContext(ctx, getTag, arg.ID, arg.OrganisationID)
	var i Tag
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Color,
		&i.OrganisationID,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listTags = `-- name: ListTags :many
SELECT id, name, color, organisation_id, created_at, deleted_at
FROM tags
WHERE organisation_id = $1
LIMIT $2 OFFSET $3
`

type ListTagsParams struct {
	OrganisationID string `json:"organisation_id"`
	Limit          int32  `json:"limit"`
	Offset         int32  `json:"offset"`
}

func (q *Queries) ListTags(ctx context.Context, arg ListTagsParams) ([]Tag, error) {
	rows, err := q.db.QueryContext(ctx, listTags, arg.OrganisationID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tag
	for rows.Next() {
		var i Tag
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Color,
			&i.OrganisationID,
			&i.CreatedAt,
			&i.DeletedAt,
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
