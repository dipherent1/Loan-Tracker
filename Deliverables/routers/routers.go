package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var Router *gin.Engine
var DataBase *mongo.Database

func Setuprouter(client *mongo.Client) *gin.Engine {
	// Initialize the database
	DataBase = client.Database("Loan-Tracker")
	fmt.Print(DataBase)

	// Initialize the router
	Router = gin.Default()

	// Initialize the Auth routes
	AuthRouter()

	return Router
}
