package Config

import (
	"log"
	"os"
)

var Port = ":8080"
var BASE_URL = "http://localhost" + Port

// Global variable to store the Env variables
var JwtSecret = []byte("your_jwt_secret")
var MONGO_CONNECTION_STRING string
// var Mail_TRAP_API_KEY string
// var GROQ_API_KEY string
// var GOOGLE_KEY string
// var GOOGLE_SECRET string
// var Google_Callback string
// var Cloud_api_key string
// var Cloud_api_secret string

func Envinit() {

	// Read JWT from environment
	JwtSecretKey := os.Getenv("JWT_SECRETE_KEY")
	if JwtSecretKey != "" {
		JwtSecret = []byte(JwtSecretKey)
	} else {
		JwtSecret = []byte("JwtSecretKey")
		log.Fatal("JWT secret key not configured")
	}
	// Read MONGO_CONNECTION_STRING from environment
	MONGO_CONNECTION_STRING = os.Getenv("MONGO_CONNECTION_STRING")
	if MONGO_CONNECTION_STRING == "" {
		MONGO_CONNECTION_STRING = "tst"
		log.Fatal("MONGO_CONNECTION_STRING is not set")
	}

}