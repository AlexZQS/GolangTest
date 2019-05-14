package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8899",
	}

	http.HandleFunc("/", welcome)
	http.HandleFunc("/doCookie", handleDoCookie)
	http.HandleFunc("/abc/efg", handleAbc)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	server.ListenAndServe()
}

func handleAbc(writer http.ResponseWriter, request *http.Request) {
	cookies := request.Cookies()
	fmt.Fprintln(writer, cookies)
}

func handleDoCookie(writer http.ResponseWriter, request *http.Request) {

	//cookie := http.Cookie{Name: "myName", Value: "myValues",HttpOnly:true}
	//cookie := http.Cookie{Name: "myName", Value: "myValues",Path:"/abc/"}
	//cookie := http.Cookie{Name: "myName", Value: "myValues", MaxAge: 10}
	//cookie := http.Cookie{Name: "myName", Value: "myValues", Expires: time.Date(2019,1,1,1,1,1,1,time.Local)}
	cookie := http.Cookie{
		Name: "myName", Value: "myValues", Domain: "www.baidu.com",
	}
	http.SetCookie(writer, &cookie)
	files, _ := template.ParseFiles("templates/cookie_set.html")
	files.Execute(writer, nil)
}

func welcome(writer http.ResponseWriter, request *http.Request) {
	files, _ := template.ParseFiles("templates/cookie_set.html")
	files.Execute(writer, nil)

}
