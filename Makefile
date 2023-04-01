gen:
	go run github.com/99designs/gqlgen generate

fmt:
	gofmt -s -l -w .