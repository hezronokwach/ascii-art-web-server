package handlers

import (
	"fmt"
	"html/template"
	"os"
)

var temp *template.Template

func init() {
	temp1, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Printf("error parsing the template %s", err)
		os.Exit(1)
	}
	temp = temp1
}
func GetTemplate() *template.Template {
	return temp
}
