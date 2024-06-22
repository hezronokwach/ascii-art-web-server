package handlers

import (
	"fmt"
	"html/template"
	"os"
)

var template *template.Template

func init() {
	temp1, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Printf("error parsing the template %s", err)
		os.Exit(1)
	}
	template = temp1
}
func GetTemplate() *template.Template {
	return template
}
