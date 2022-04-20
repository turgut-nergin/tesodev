package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/routes"
	"github.com/turgut-nergin/tesodev/mongo"
	repositorty "github.com/turgut-nergin/tesodev/repository/repo"
)

func main() {
	conn := os.Getenv("mongodb://mongo-db:27017")

	mongoClient, err := mongo.NewClient(conn)
	if err != nil {
		panic("Connection could not be established")
	}
	repo := repositorty.New(mongoClient)
	engine := gin.New()
	engine.Use(gin.Recovery())
	routes.InitializeRoutes(engine, repo)
	engine.Run(":8086")
}
