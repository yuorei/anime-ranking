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
	"github.com/yuorei/anime-ranking/graph"
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

	// srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// // リクエストの認証に必要なJWTトークンの検証を行うauthMiddleware関数を追加する
	// srv.Use(middlewares.AuthMiddleware)
	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	router := mux.NewRouter()
	router.Use(middlewares.AuthMiddleware)

	c := graph.Config{Resolvers: &graph.Resolver{}}
	c.Directives.Auth = directives.Auth

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))
	// CORSの設定
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // 許可するOriginを指定
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowedHeaders: []string{"Content-Type"},
	})
	router.PathPrefix("/query").Handler(corsOpts.Handler(srv))
	router.PathPrefix("/").Handler(playground.Handler("GraphQL playground", "/query"))

	// router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
