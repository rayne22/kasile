package routes

import (
	"github.com/gin-gonic/gin"
	"kasile/controllers"
)

// Serve keeps gin framework instance
type Serve struct {
	Engine *gin.Engine
}

func (e *Serve) Routes() *Serve {

	// Loads all the templates
	e.Engine.LoadHTMLGlob("templates/**/*")

	// Home Route
	home := e.Engine.Group("/")
	{
		home.GET("", controllers.IndexController)
		home.GET("test", controllers.TestController)
	}

	e.Engine.Static("/assets", "./static/assets")

	//e.Engine.StaticFS("/more_static", http.Dir("my_file_system"))
	//e.Engine.StaticFile("/favicon.ico", "./static/assets/favicon.ico")

	return e
}
