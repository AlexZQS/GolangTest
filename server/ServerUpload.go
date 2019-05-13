package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

/**
 * @description
 * @time 2019/5/13 23:31
 * @version
 */
func main() {
	server := http.Server{Addr: ":8070"}
	http.HandleFunc("/server/uploadPage", ServerUploadPage)
	http.HandleFunc("/server/upload", ServerUpload)
	server.ListenAndServe()
}

func ServerUpload(writer http.ResponseWriter, request *http.Request) {
	fileName := request.FormValue("name")
	file, fileHeader, _ := request.FormFile("file")
	bytes, _ := ioutil.ReadAll(file)
	fileName = fileName + fileHeader.Filename[strings.LastIndex(fileHeader.Filename, "."):]
	ioutil.WriteFile("G:/"+fileName, bytes, 0777)
	files, _ := template.ParseFiles("templates/upload_success.html")
	files.Execute(writer, nil)

}

func ServerUploadPage(writer http.ResponseWriter, request *http.Request) {
	files, _ := template.ParseFiles("templates/upload.html")
	files.Execute(writer, nil)
}
