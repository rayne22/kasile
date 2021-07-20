package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// IndexController Runs index page
var AboutController = func(c *gin.Context) {

	c.HTML(http.StatusOK, "about-page", gin.H{
		"title": "Kasile",
	})

}
