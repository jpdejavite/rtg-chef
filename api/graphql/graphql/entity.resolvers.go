package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	generated1 "github.com/jpdejavite/rtg-chef/api/graphql/generated"
	generated "github.com/jpdejavite/rtg-chef/api/graphql/graph/model"
	"github.com/jpdejavite/rtg-chef/api/graphql/models"
)

func (r *entityResolver) FindAppQueriesByID(ctx context.Context, id string) (*generated.AppQueries, error) {
	return &generated.AppQueries{Recipes: &models.RecipeAppQueries{}}, nil
}

// Entity returns generated1.EntityResolver implementation.
func (r *Resolver) Entity() generated1.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
