package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	server := http.Server{Addr: ":8899"}
	http.HandleFunc("/param", handleParam)
	server.ListenAndServe()
}

func handleParam(writer http.ResponseWriter, request *http.Request) {
	headers := request.Header
	fmt.Fprintln(writer, headers["Accept-Encoding"])
	fmt.Fprintln(writer, len(headers["Accept-Encoding"]))
	for _, n := range strings.Split(headers["Accept-Encoding"][0], ",") {
		fmt.Fprintln(writer, strings.TrimSpace(n))
	}

	//必须先解析成Form
	request.ParseForm()
	fmt.Fprintln(writer, request.Form)
	fmt.Fprintln(writer, request.Form["name"][0])

	//第二种取值
	fmt.Fprintln(writer, request.FormValue("name"))
}
