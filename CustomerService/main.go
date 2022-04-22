package main

import (
	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/routes"
	"github.com/turgut-nergin/tesodev/middleware"
	"github.com/turgut-nergin/tesodev/mongo"
	repositorty "github.com/turgut-nergin/tesodev/repository/repo"
)

func main() {
	// client := mongo.GetMongoDB()

	client := mongo.NewClient()
	repo := repositorty.New(client)

	engine := gin.New()
	engine.Use(middleware.CORSMiddleware())
	engine.Use(gin.Recovery())
	routes.InitializeRoutes(engine, repo)
	engine.Run(":8086")
}