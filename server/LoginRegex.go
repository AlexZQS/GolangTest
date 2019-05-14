package main

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
)

func main() {
	server := http.Server{
		Addr: ":8899"}

	http.HandleFunc("/loginPage", loginRegexPage)
	http.HandleFunc("/login", loginRegex)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	server.ListenAndServe()
}

func loginRegexPage(writer http.ResponseWriter, request *http.Request) {
	files, _ := template.ParseFiles("templates/login_regex_match.html")
	files.Execute(writer, nil)
}

func loginRegex(writer http.ResponseWriter, request *http.Request) {

	username := request.FormValue("username")
	matched, _ := regexp.MatchString(`^[0-9a-zA-Z]{6,12}$`, username)
	if matched {
		fmt.Fprintln(writer, "登录成功，数据库已经插入")
	} else {
		fmt.Fprintln(writer, "格式不正确")
	}
}
