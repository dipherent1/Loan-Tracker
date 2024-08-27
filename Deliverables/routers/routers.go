package routers

import (
	"fmt"
	custommongo "loaner/CustomMongo"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var DataBase custommongo.Database


func Setuprouter(client custommongo.Client) *gin.Engine {
	// Initialize the database
	DataBase = client.Database("Loan-Tracker")
	fmt.Print(DataBase)

	// Initialize the router
	router = gin.Default()

	// Initialize the Auth routes
	AuthRouter()

	return router
}
