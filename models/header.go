package models

import (
	"html/template"
	"log"
	"net/http"
)

// Header keeps header information
type Header struct {
	SectionID string
	SectionTitle string
	Title string
	Logo string
}

func LoadIndex(w http.ResponseWriter)  {

	var headers = []Header {
		{
			SectionID:    "#hero",
			SectionTitle: "Header",
		},
		{
			SectionID:    "#about",
			SectionTitle: "About",
		},
		{
			SectionID:    "#services",
			SectionTitle: "Services",
		},
		{
			SectionID:    "#portfolio",
			SectionTitle: "Portfolio",
		},
		{
			SectionID:    "#team",
			SectionTitle: "Team",
		},
		{
			SectionID:    "#contact",
			SectionTitle: "Contact",
		},
		//{
		//	SectionID:    "#about",
		//	SectionTitle: "Get Started",
		//},
	}


	tpl := template.Must(template.ParseGlob("templates/*/*.gohtml"))



	// passes the template using the path
	err := tpl.ExecuteTemplate(w, "index" , headers)


	if err != nil {
		log.Println("EE",err)
		return
	}
}


func MkSlice(args ...interface{}) []interface{} {
	return args
}
