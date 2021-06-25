package serve

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"kasile/routes"
	"log"
	"os"
)

func Serve() {
	//Loads .env file
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}

	// Sets the mode of the Web (Either Release or Debug mode)
	gin.SetMode(gin.DebugMode)

	// Defines the Gin Gonic Engine
	router := gin.Default()

	// Instantiates Serve struct
	route := routes.Serve{}

	route.Engine = router

	_ = route.Routes()

	// Instantiates the Port the Web App runs on
	port := fmt.Sprintf("%s%s", ":", os.Getenv("PORT"))

	// Starts listening and serving
	_ = route.Engine.Run(port)

}
