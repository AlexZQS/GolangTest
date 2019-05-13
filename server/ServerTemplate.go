package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	server := http.Server{Addr: ":8090"}
	http.HandleFunc("/templates", handleTemplates)
	server.ListenAndServe()
}

func handleTemplates(writer http.ResponseWriter, request *http.Request) {
	files, err := template.ParseFiles("templates/Template_start.html")
	if err != nil {
		fmt.Println("解析出错")
		return
	}
	files.Execute(writer, nil)
}
