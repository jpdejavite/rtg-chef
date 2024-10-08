GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
DOCKER_BUILD=$(shell pwd)/.docker_build
DOCKER_CMD=$(DOCKER_BUILD)/go-getting-started

$(DOCKER_CMD): clean
	mkdir -p $(DOCKER_BUILD)
	$(GO_BUILD_ENV) go build -v -o $(DOCKER_CMD) cmd/graphqlServer.go

clean:
	rm -rf $(DOCKER_BUILD)

heroku: $(DOCKER_CMD)
	heroku container:push web

gen-gql:
	cd api/graphql
	go run github.com/99designs/gqlgen generate --verbose
	cd ../../

test :
	go clean -testcache
	go test ./... -race -coverprofile cp.out
	go tool cover -html=./cp.out -o cover.html

run-dev:
	make -f cmd/dev/Makefile

run:
	go build -o bin/rtg-chef -v cmd/graphqlServer.go
	heroku local web
