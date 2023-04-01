package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"github.com/yuorei/anime-ranking/database/mysql"
	"github.com/yuorei/anime-ranking/directives"
	"github.com/yuorei/anime-ranking/graph/generated"
	"github.com/yuorei/anime-ranking/graph/resolver"
	"github.com/yuorei/anime-ranking/middlewares"
)

const defaultPort = "8080"

func main() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	mysql.NewMySQL()
	mysql.CreateTable()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := mux.NewRouter()
	router.Use(middlewares.AuthMiddleware)

	c := generated.Config{Resolvers: &resolver.Resolver{}}
	c.Directives.Auth = directives.Auth

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))
	// CORSの設定
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // 許可するOriginを指定
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowedHeaders: []string{"Content-Type"},
	})
	router.PathPrefix("/query").Handler(corsOpts.Handler(srv))
	router.PathPrefix("/").Handler(playground.Handler("GraphQL playground", "/query"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
