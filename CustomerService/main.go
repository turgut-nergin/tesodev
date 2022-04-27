package main

import (
	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/routes"
	"github.com/turgut-nergin/tesodev/config"
	repository "github.com/turgut-nergin/tesodev/database"
	"github.com/turgut-nergin/tesodev/database/models"
	"github.com/turgut-nergin/tesodev/mongo"
)

func main() {
	// client := mongo.GetMongoDB()
	dbModel := models.Repository{
		Name:           "tesodev",
		CollectionName: "customers",
	}

	url := "mongo-db:27017"
	client := mongo.NewClient(url)
	repo := repository.New(client, dbModel)

	engine := gin.New()
	engine.Use(config.InitCORSConfig())
	engine.Use(gin.Recovery())
	routes.GetRouter(engine, repo)
	engine.Run(":8086")
}
