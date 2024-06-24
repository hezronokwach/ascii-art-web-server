package handlers

import (
	"fmt"
	"net/http"

	//"html/template"
	asciiart "asciiart/functionFiles"
)

const (
	notFound            = http.StatusNotFound
	internalServerError = http.StatusInternalServerError
	badRequest          = http.StatusBadRequest
	methodNotAllowed    = http.StatusMethodNotAllowed
)

func Request(writer http.ResponseWriter, reader *http.Request) {
	if reader.URL.Path != "/" {
		msg := fmt.Sprintf("%d% s", notFound, " Page not found")
		http.Error(writer, msg, http.StatusNotFound)
		return
	}
	if reader.Method != http.MethodGet {
		msg := fmt.Sprintf("%d% s", methodNotAllowed, " Method not allowed")
		http.Error(writer, msg, http.StatusMethodNotAllowed)
		return
	}
	temp1 := GetTemplate()
	temp1.Execute(writer, Data{Success: false})
}

func Post(writer http.ResponseWriter, reader *http.Request) {
	temp2 := GetTemplate()
	if reader.Method != http.MethodPost {
		msg := fmt.Sprintf("%d% s", methodNotAllowed, " Method not allowed")
		http.Error(writer, msg, http.StatusMethodNotAllowed)
		return
	}
	userInput := reader.FormValue("userinput")
	banner := reader.FormValue("bannerfile")

	characteMap, err := asciiart.CreateMap(banner)
	if err != nil {
		msg := fmt.Sprintf("%d% s", internalServerError, " Internal Server Error ")
		http.Error(writer, msg, http.StatusInternalServerError)
		fmt.Printf("Error creating map %s", err)
		fmt.Println()
		return
	}
	result := asciiart.DisplayAsciiArt(characteMap, userInput)
	if result == "" {
		msg := fmt.Sprintf("%d% s", badRequest, " Bad Request ")
		http.Error(writer, msg, http.StatusBadRequest)
		fmt.Printf("Character not found %s", err)
		fmt.Println()
		return
	}
	temp2.Execute(writer, Data{Success: true, Result: result})
}
