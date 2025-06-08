package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	hackernews "github.com/sharukh010/hackernews/graph"
	database "github.com/sharukh010/hackernews/internal/pkg/db/migrations/mysql"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == ""{
		port = defaultPort
	}

	router := chi.NewRouter()

	database.InitDB()
	defer database.CloseDB()
	database.Migrate()
	
	server := handler.NewDefaultServer(hackernews.NewExecutableSchema(hackernews.Config{Resolvers:&hackernews.Resolver{}}))
	router.Handle("/",playground.Handler("GraphQL playground","/query"))
	router.Handle("/query",server)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground\n",port)
	log.Fatal(http.ListenAndServe(":"+port,router))
}
