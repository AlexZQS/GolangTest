package main

import (
	"html/template"
	"net/http"
	"time"
)

func main() {

	server := http.Server{
		Addr: ":8091",
	}

	http.HandleFunc("/templateFunc", handleTemplateFunc)
	server.ListenAndServe()
}

func MyTransfer(t time.Time) string {
	return t.Format("2005-01-04 15:03:00")
}

func handleTemplateFunc(writer http.ResponseWriter, request *http.Request) {
	fm := template.FuncMap{"mt": MyTransfer}
	t := template.New("template_func.html").Funcs(fm)
	t, _ = t.ParseFiles("templates/template_func.html")

	//files, _ := template.ParseFiles("templates/template_func.html")
	dateTime := time.Date(2019, 1, 2, 3, 4, 5, 0, time.Local)
	t.Execute(writer, dateTime)

}
