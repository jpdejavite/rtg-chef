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
	"github.com/jpdejavite/rtg-chef/pkg/gql-base-lib/config"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

// ValidateAuthorize proccess authorize directive and check resource roles against token roles
func ValidateAuthorize(ctx context.Context, roles []*string) error {

	// 	reqCred := ctx.Value("requestedCredential").(RequestedCredential)

	// 	hasAnyRole := false
	// 	userRoles := reqCred.Roles

	// 	for _, reqRole := range roles {
	// 		for _, userRole := range userRoles {
	// 			if reqRole != nil && *reqRole == userRole {
	// 				hasAnyRole = true
	// 				break
	// 			}
	// 		}
	// 	}

	// 	if !hasAnyRole {
	// 		return errors.New(AccessDeniedError.Code, AccessDeniedError.Message)
	// 	}

	return nil
}

// /*AddSecurityHandler extracts security credentials sent by gateway from header
// and put into request context app using a standard struct */
// func AddSecurityHandler(authCfg Config) func(next http.Handler) http.Handler {

// 	addCtx := func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 			if r.Header.Get("gateway-authorization") == "" {
// 				next.ServeHTTP(w, r)
// 				return
// 			}

// 			authCfg.Token = r.Header.Get("gateway-authorization")

// 			headers, err := ValidateToken(authCfg)

// 			authCfg.TokenVersion = headers["version"].(string)

// 			if err != nil {
// 				smlog.Error("server.go", "AddSecurityHandler", err.Error(), "")
// 				next.ServeHTTP(w, r)
// 				return
// 			}

// 			reqCred := RequestedCredential{}
// 			var ctx context.Context

// 			if mapOfClaims, ok := headers["claims"]; ok == true {
// 				claims := mapOfClaims.(map[string]interface{})

// 				if claims["uid"] != nil {
// 					uid := claims["uid"].(string)
// 					reqCred.UserID = &uid
// 				}

// 				clientName := headers["iss"].(string)
// 				reqCred.ClientName = &clientName

// 				roles := func() []interface{} {
// 					if claims["roles"] != nil {
// 						return claims["roles"].([]interface{})
// 					}
// 					if claims["role"] != nil {
// 						return claims["role"].([]interface{})
// 					}
// 					return []interface{}{}
// 				}()

// 				for _, rol := range roles {
// 					reqCred.Roles = append(reqCred.Roles, rol.(string))
// 				}
// 				ctx = context.WithValue(r.Context(), "requestedCredential", reqCred)
// 			}

// 			next.ServeHTTP(w, r.WithContext(ctx))
// 		})
// 	}

// 	return addCtx
// }

func configServer() generated.Config {

	directive := func(ctx context.Context, obj interface{}, next gql.Resolver, roles []*string) (res interface{}, err error) {
		err = ValidateAuthorize(ctx, roles)

		if err != nil {
			return nil, err
		}
		return next(ctx)
	}

	c := generated.Config{Resolvers: &gr.Resolver{}, Directives: generated.DirectiveRoot{HasAllRoles: directive}}

	return c
}

func Start() {
	coi := log.GenerateCoi(nil)
	config.LoadAll(constants.GetEnVarKeys())
	port := fmt.Sprintf(":%s", config.GetEnvVar(constants.Port))

	cfg := configServer()

	router := chi.NewRouter()
	// authCfg := auth.Config{
	// 	SeverinoBasicToken: fmt.Sprintf("Basic %s", jwt.CreateSeverinoToken()),
	// 	SeverinoURL:        fmt.Sprintf("%s/config", viper.GetString("SEVERINO_URL"))}

	// router.Use(AddSecurityHandler(authCfg))
	serv := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))

	serv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		// if smErr, ok := err.(errors.ErrorWrapper); ok {
		// 	gqlErr := &gqlerror.Error{
		// 		Message:    smErr.Message,
		// 		Extensions: map[string]interface{}{"code": smErr.Code},
		// 	}
		// 	return gqlErr
		// }

		return graphql.DefaultErrorPresenter(ctx, err)
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", serv)

	message := fmt.Sprintf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Info("server", message, nil, coi)
	if err := http.ListenAndServe(port, router); err != nil {
		panic(err)
	}

}
