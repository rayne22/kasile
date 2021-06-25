package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// IndexController Runs index page
var IndexController = func(c *gin.Context) {

	c.HTML(http.StatusOK, "index", gin.H{
		"title": "Test",
	})

}
