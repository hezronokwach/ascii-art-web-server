package handlers

import (
	"fmt"
	"net/http"
	"os"

	//"html/template"
	asciiart "asciiart/functionFiles"
)

func Request(writer http.ResponseWriter, reader *http.Request) {
	temp1 := GetTemplate()
	if reader.Method != http.MethodGet {
		http.Error(writer, "405 Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	temp1.Execute(writer, Data{Success: false})
}

func Post(writer http.ResponseWriter, reader *http.Request) {
	temp2 := GetTemplate()
	if reader.Method != http.MethodPost {
		http.Error(writer, "405 Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	userInput := reader.FormValue("userinput")
	banner := reader.FormValue("bannerfile")

	characteMap, err := asciiart.CreateMap(banner)
	if err != nil {
		fmt.Printf("Error creating map %s", err)
		fmt.Println()
		os.Exit(1)
	}
	result := asciiart.DisplayAsciiArt(characteMap, userInput)
	temp2.Execute(writer, Data{Success:true, Result:result})
	fmt.Println(userInput)
	fmt.Println(banner)
	
}
