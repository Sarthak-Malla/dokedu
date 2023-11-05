package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"errors"
	"example/internal/db"
	"example/internal/graph/model"
	"example/internal/middleware"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// UpdateOrganisation is the resolver for the updateOrganisation field.
func (r *mutationResolver) UpdateOrganisation(ctx context.Context, input model.UpdateOrganisationInput) (*db.Organisation, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}
	if !currentUser.HasPermissionAdmin() {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "You do not have permission to update this organisation",
			Extensions: map[string]interface{}{
				"code": "UNAUTHORIZED",
			},
		})
		return nil, errors.New("unauthorized")
	}
	if currentUser.OrganisationID != input.ID {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "You do not have permission to update this organisation",
			Extensions: map[string]interface{}{
				"code": "UNAUTHORIZED",
			},
		})
		return nil, errors.New("unauthorized")
	}

	var organisation db.Organisation
	err = r.DB.NewSelect().Model(&organisation).Where("id = ?", currentUser.OrganisationID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	if input.Name != nil {
		organisation.Name = *input.Name
	}

	if input.LegalName != nil {
		organisation.LegalName = *input.LegalName
	}

	if input.Phone != nil {
		organisation.Phone = *input.Phone
	}

	if input.Website != nil {
		organisation.Website = *input.Website
	}

	err = r.DB.NewUpdate().Model(&organisation).Column("name").Column("legal_name").Column("phone").Column("website").WherePK().Returning("*").Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &organisation, nil
}

// Owner is the resolver for the owner field.
func (r *organisationResolver) Owner(ctx context.Context, obj *db.Organisation) (*db.User, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	var user db.User
	err = r.DB.NewSelect().Model(&user).Where("id = ?", obj.OwnerID).Where("organisation_id = ?", currentUser.OrganisationID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Organisation is the resolver for the organisation field.
func (r *queryResolver) Organisation(ctx context.Context) (*db.Organisation, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	var organisation db.Organisation
	err = r.DB.NewSelect().Model(&organisation).Where("id = ?", currentUser.OrganisationID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &organisation, nil
}

// Organisation returns OrganisationResolver implementation.
func (r *Resolver) Organisation() OrganisationResolver { return &organisationResolver{r} }

type organisationResolver struct{ *Resolver }
