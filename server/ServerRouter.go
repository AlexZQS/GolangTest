package main

import "net/http"

type MyHandle struct {
}

type MyHandler struct {
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("MyHandler的数据 --第二个"))
}

func (m *MyHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("返回的数据 -- 第一个"))
}

func main() {
	h := MyHandle{}
	h2 := MyHandler{}
	server := http.Server{
		Addr: "localhost:8889",
	}

	http.Handle("/first", &h)
	http.Handle("/second", &h2)

	server.ListenAndServe()
}
