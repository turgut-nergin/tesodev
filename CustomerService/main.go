package main

import (
	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/routes"
	repository "github.com/turgut-nergin/tesodev/database"
	"github.com/turgut-nergin/tesodev/middleware"
	"github.com/turgut-nergin/tesodev/mongo"
)

func main() {
	// client := mongo.GetMongoDB()

	client := mongo.NewClient()
	repo := repository.New(client)

	engine := gin.New()
	engine.Use(middleware.CORSMiddleware())
	engine.Use(gin.Recovery())
	routes.InitializeRoutes(engine, repo)
	engine.Run(":8086")
}
