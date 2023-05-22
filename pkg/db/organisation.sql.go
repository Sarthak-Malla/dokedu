// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: organisation.sql

package db

import (
	"context"

	"github.com/lib/pq"
)

const getOrganisation = `-- name: GetOrganisation :one
SELECT id, name, legal_name, website, phone, owner_id, allowed_domains, created_at, deleted_at
FROM organisations
WHERE id = $1
`

func (q *Queries) GetOrganisation(ctx context.Context, id string) (Organisation, error) {
	row := q.db.QueryRowContext(ctx, getOrganisation, id)
	var i Organisation
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LegalName,
		&i.Website,
		&i.Phone,
		&i.OwnerID,
		pq.Array(&i.AllowedDomains),
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}