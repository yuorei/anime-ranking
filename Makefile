gen:
	go run github.com/99designs/gqlgen generate

fmt:
	gofmt -s -l -w .
build:
	@docker compose build
up:
	@docker compose up

ps:
	@docker compose ps

stop:
	@docker compose stop
