package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg/v9"
	"github.com/rs/cors"
	"github.com/sony-nurdianto/scratch-go/graph"
	"github.com/sony-nurdianto/scratch-go/graph/domain"
	"github.com/sony-nurdianto/scratch-go/graph/generated"
	"github.com/sony-nurdianto/scratch-go/graph/middleware1"
	"github.com/sony-nurdianto/scratch-go/graph/postgres"
)

const defaultPort = "8080"

func main() {

	r := chi.NewRouter()

	DB := postgres.New(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "scratch-go",
	})

	defer DB.Close()

	DB.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	usersRepo := postgres.UsersRepo{DB: DB}

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware1.AuthMiddleware(usersRepo))

	d := domain.NewDomain(usersRepo, postgres.BucketRepo{DB: DB}, postgres.ProductRepo{DB: DB})

	c := generated.Config{Resolvers: &graph.Resolver{Domain: d}}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
