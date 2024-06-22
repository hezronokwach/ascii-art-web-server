package handlers

import (
	//"html/template"
	"net/http"
)

func Request(writer http.ResponseWriter , reader * http.Request) {
	temp1:=GetTemplate()
	if reader.Method != http.MethodGet{
		http.Error(writer, "405 Invalid request method", http.StatusMethodNotAllowed )
		return
	}
	temp1.Execute(writer, Data{Success: false})
}
