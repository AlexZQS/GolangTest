package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

func main() {
	server := http.Server{
		Addr: ":8899"}
	http.HandleFunc("/", handleJson)
	http.HandleFunc("/showUser", handleShowUser)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	server.ListenAndServe()
}

func handleShowUser(writer http.ResponseWriter, request *http.Request) {
	users := make([]User, 0)
	users = append(users, User{"张三", 17})
	users = append(users, User{"李四", 16})
	users = append(users, User{"王五", 19})
	writer.Header().Set("Content-Type", "application/json;Charset=utf-8")
	bytes, _ := json.Marshal(users)
	fmt.Fprintln(writer, string(bytes))
}

func handleJson(writer http.ResponseWriter, request *http.Request) {
	files, _ := template.ParseFiles("templates/parse_json.html")
	files.Execute(writer, nil)
}
