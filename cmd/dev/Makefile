.EXPORT_ALL_VARIABLES:

include cmd/dev/dev.env

run-dev:
	go build -o bin/rtg-chef -v cmd/graphqlServer.go
	./bin/rtg-chef
