package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/alex/{url}", hello)
	router.HandleFunc("/test", test)

	http.ListenAndServe(":8090", router)
}

func test(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "这是Test")
}

func hello(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintln(writer, vars["url"])
}
