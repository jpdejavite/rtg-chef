package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	generated1 "github.com/jpdejavite/rtg-chef/api/graphql/generated"
	generated "github.com/jpdejavite/rtg-chef/api/graphql/graph/model"
	"github.com/jpdejavite/rtg-chef/api/graphql/models"
	"github.com/jpdejavite/rtg-chef/internal/chef/constants"
	"github.com/jpdejavite/rtg-go-toolkit/pkg/config"
	"github.com/jpdejavite/rtg-go-toolkit/pkg/graphql/auth"
)

func (r *recipeAppQueriesResolver) List(ctx context.Context, obj *models.RecipeAppQueries, input generated.RecipeListInput) (*generated.RecipeList, error) {
	authData, gc, c, coi := auth.GetContextInfo(ctx)
	fmt.Println("authData", authData)
	fmt.Println("gc", (*gc).GetGlobalConfigAsStr(config.GatewayPublicKey))
	fmt.Println("c", (*c).GetConfigAsStr(constants.DatabaseURL))
	fmt.Println("coi", coi)
	return &generated.RecipeList{
		Total: 2,
		Recipes: []*generated.Recipe{
			&generated.Recipe{Name: "Receita 1", Description: "Descrição 1"},
			&generated.Recipe{Name: "Receita 2", Description: "Descrição 2"},
		},
	}, nil
}

// RecipeAppQueries returns generated1.RecipeAppQueriesResolver implementation.
func (r *Resolver) RecipeAppQueries() generated1.RecipeAppQueriesResolver {
	return &recipeAppQueriesResolver{r}
}

type recipeAppQueriesResolver struct{ *Resolver }
