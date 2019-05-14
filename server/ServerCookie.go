package main

import (
	"html/template"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8899",
	}

	http.HandleFunc("/cookiePage", handleCookiePage)
	http.HandleFunc("/setCookie", handleSetCookie)
	http.HandleFunc("/getCookie", handleGetCookie)
	server.ListenAndServe()
}

func handleGetCookie(writer http.ResponseWriter, request *http.Request) {
	//根据key取出cookie
	//cookie, _ := request.Cookie("myKey")
	//取出全部cookie
	cookies := request.Cookies()
	files, _ := template.ParseFiles("templates/cookie.html")
	files.Execute(writer, cookies)
}

func handleCookiePage(writer http.ResponseWriter, request *http.Request) {
	files, _ := template.ParseFiles("templates/cookie.html")
	files.Execute(writer, nil)

}

func handleSetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := http.Cookie{
		Name:  "myKey",
		Value: "myValue",
	}
	http.SetCookie(writer, &cookie)
	files, _ := template.ParseFiles("templates/cookie.html")
	files.Execute(writer, nil)
}
