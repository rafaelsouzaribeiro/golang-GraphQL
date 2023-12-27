package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rafaelsouzaribeiro/13-graphql/graph"
	"github.com/rafaelsouzaribeiro/13-graphql/internal/database"
)

const defaultPort = "8080"

func main() {
	db, err := sql.Open("sqlite3", "./data.db")

	defer db.Close()
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	categoryDb := database.NewCategory(db)
	courseDb := database.NewCouser(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CategoryDB: categoryDb,
		CouseDB:    courseDb,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
