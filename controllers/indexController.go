package controllers

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"kasile/models"
	"log"
	"net/http"
)

// IndexController Runs index page
var IndexController = func(c *gin.Context) {
	//var h models.Header
	//h.Title = "Test"
	//tpl := template.Must(template.ParseGlob("templates/*/*.gohtml"))
	//
	//
	//h.Title = "MEE"
	//
	//// passes the template using the path
	//
	////var body bytes.Buffer
	//
	//err := tpl.ExecuteTemplate(c.Writer, "index" , h)
	//
	//
	//if err != nil {
	//	log.Println("EE",err)
	//	return
	//}
	//models.LoadIndex(c.Writer)
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "Kasile",
	})

}

var TestController = func(c *gin.Context) {

	h := models.Header{}

	//var h Header
	path := "templates/layouts/test.gohtml"
	h.Title = "Test"
	// passes the template using the path
	t := template.New("Test")

	t, err := template.ParseFiles(path)

	if err != nil {
		log.Println(err)
		return
	}

	//var body bytes.Buffer

	err = t.Execute(c.Writer, h)

	log.Println(h.Title, "dasd")

	if err != nil {
		log.Println("EE", err)
		return
	}

	//models.LoadIndex()
	c.HTML(http.StatusOK, "test", gin.H{
		"title": "Test",
	})

}
