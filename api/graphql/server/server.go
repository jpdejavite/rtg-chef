package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	gql "github.com/99designs/gqlgen/graphql"
	"github.com/jpdejavite/go-log/pkg/log"
	"github.com/jpdejavite/rtg-chef/api/graphql/generated"
	gr "github.com/jpdejavite/rtg-chef/api/graphql/graphql"
	"github.com/jpdejavite/rtg-chef/internal/chef/constants"
	"github.com/jpdejavite/rtg-go-toolkit/pkg/config"
	"github.com/jpdejavite/rtg-go-toolkit/pkg/envvar"
	"github.com/jpdejavite/rtg-go-toolkit/pkg/firestore"
	"github.com/jpdejavite/rtg-go-toolkit/pkg/graphql/auth"
	"github.com/jpdejavite/rtg-go-toolkit/pkg/graphql/errors"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

func configServer() generated.Config {

	directive := func(ctx context.Context, obj interface{}, next gql.Resolver, roles []*string) (res interface{}, err error) {
		err = auth.ValidateHasAllRoles(ctx, roles)

		if err != nil {
			return nil, err
		}
		return next(ctx)
	}

	c := generated.Config{Resolvers: &gr.Resolver{}, Directives: generated.DirectiveRoot{HasAllRoles: directive}}

	return c
}

// Start start server
func Start() {
	coi := log.GenerateCoi(nil)
	envvar.LoadAll(constants.GetEnVarKeys())
	port := fmt.Sprintf(":%s", envvar.GetEnvVar(constants.Port))

	if err := firestore.ConnectToDatabase(envvar.GetEnvVar(constants.FirebaseCredential)); err != nil {
		log.Fatal("server", "cannot connect to database", nil, coi)
		panic(err)
	}

	if err := config.LoadGlobalConfig(); err != nil {
		log.Fatal("server", "error loading global config", nil, coi)
		panic(err)
	}

	cfg := configServer()

	router := chi.NewRouter()
	router.Use(auth.AddSecurityHandler())
	serv := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))

	serv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		if customError, ok := err.(errors.CustomError); ok {
			gqlError := &gqlerror.Error{
				Message:    customError.Message,
				Extensions: map[string]interface{}{"code": customError.Code},
			}
			return gqlError
		}

		return graphql.DefaultErrorPresenter(ctx, err)
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", serv)

	message := fmt.Sprintf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Info("server", message, nil, coi)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("server", "cannot listen to port", map[string]string{
			"error": err.Error(),
		}, coi)
		panic(err)
	}

}
