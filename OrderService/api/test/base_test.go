package test

import (
	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/handlers/order"
	"github.com/turgut-nergin/tesodev/database"
	"github.com/turgut-nergin/tesodev/database/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	// Init API
	api := gin.Default()

	// Using Middle-wares get the db session in the context for each request
	api.Use(middlewares.Connect)

	var a *database.Repository
	routerProjects := api.Group("/projects")
	{
		// All Projects
		routerProjects.GET("", order.GetOrdersHandler(a))
	}

}

func ListAllProjects(c *gin.Context) {

	session := c.MustGet("session").(*mgo.Session)
	db := session.DB("myDB")

	var projects []models.Order
	err := db.C("projects").Find(bson.M{}).All(&projects)

	if err != nil {
		helpers.RespondWithError(400, err.Error(), c)
	}

	c.JSON(200, projects)

}
