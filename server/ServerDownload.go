package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
 * @description
 * @time 2019/5/13 23:31
 * @version
 */
func main() {
	server := http.Server{Addr: ":8070"}
	http.HandleFunc("/server/downloadPage", ServerDownloadPage)
	http.HandleFunc("/server/download", ServerDownload)
	server.ListenAndServe()
}

func ServerDownloadPage(writer http.ResponseWriter, request *http.Request) {

}

func ServerDownload(writer http.ResponseWriter, request *http.Request) {
	fileName := request.FormValue("filename")
	bytes, err := ioutil.ReadFile("D:/serverfile/" + fileName)
	if err != nil {
		fmt.Fprintln(writer, "文件下载失败", err)
		return
	}
	header := writer.Header()
	header.Set("Content-Type", "application/octet-stream")
	header.Set("Content-Disposition", "attachment;filename="+fileName)
	writer.Write(bytes)

}
