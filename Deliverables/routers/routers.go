package routers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var Router *gin.Engine
var DataBase *mongo.Database

func Setuprouter(client *mongo.Client) *gin.Engine {
	// Initialize the database
	DataBase = client.Database("Loan-Tracker")

	// Initialize the router
	Router = gin.Default()

	// Initialize the Auth routes
	AuthRouter()
	// Initialize the Refresh routes
	RefreshTokenRouter()
	// Initialize the User routes
	UserRouter()

	return Router
}
