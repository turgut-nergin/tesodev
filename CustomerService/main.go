package main

import (
	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/routes"
	"github.com/turgut-nergin/tesodev/config"
	repository "github.com/turgut-nergin/tesodev/database"
	"github.com/turgut-nergin/tesodev/mongo"
)

func main() {
	// client := mongo.GetMongoDB()

	client := mongo.NewClient()
	repo := repository.New(client)

	engine := gin.New()
	engine.Use(config.InitCORSConfig())
	engine.Use(gin.Recovery())
	routes.GetRouter(engine, repo)
	engine.Run(":8086")
}
