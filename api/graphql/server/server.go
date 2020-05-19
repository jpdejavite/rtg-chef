package server

import (
	"context"
	"fmt"
	"net/http"
	"os"

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

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"

	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
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

	firestoreDB, err := firestore.ConnectToDatabase(envvar.GetEnvVar(constants.FirebaseCredential))
	if err != nil {
		log.Fatal("server", "cannot connect to database", nil, coi)
		panic(err)
	}

	gc := config.NewGlobalConfigs(firestoreDB)
	if err := gc.LoadGlobalConfig(); err != nil {
		log.Fatal("server", "error loading global config", nil, coi)
		panic(err)
	}

	c := config.NewConfigs(firestoreDB)
	if err := c.LoadConfig(constants.AppName, constants.GetConfigKeys()); err != nil {
		log.Fatal("server", "error loading config", nil, coi)
		panic(err)
	}

	postgresDB, err := sql.Open("postgres", c.GetConfigAsStr(constants.DatabaseURL))
	driver, err := postgres.WithInstance(postgresDB, &postgres.Config{})
	if err != nil {
		log.Fatal("server", "error creating database instance", nil, coi)
		panic(err)
	}
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("server", "error getting current dir", nil, coi)
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s/scripts/migration", dir),
		"postgres", driver)
	if err != nil {
		log.Fatal("server", "error migrating database", nil, coi)
		panic(err)
	}
	m.Steps(2)

	cfg := configServer()

	router := chi.NewRouter()
	router.Use(auth.AddSecurityHandler(gc, c))
	serv := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))

	serv.SetErrorPresenter(errors.HandleGraphqlError())

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
