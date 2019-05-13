package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)

/**
 * @description
 * @time 2019/5/11 0:13
 * @version
 */

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/index", handleIndex)
	router.HandleFunc("/home", handleHome)
	router.HandleFunc("/detail", handleDetail)
	//router.HandleFunc("/db", handleDb)
	router.HandleFunc("/form", handleForm)
	router.HandleFunc("/formadd", handleFormAdd)

	//静态服务器,当发现url以/static 开头把请求转发给指定的路径
	//http://localhost:8090/static/index
	fileServer := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))

	router.HandleFunc("/setsession", handleSetSession)
	router.HandleFunc("/getsession", handleGetSession)

	http.Handle("/", router)
	_ = http.ListenAndServe(":8888", nil)
}

func handleGetSession(writer http.ResponseWriter, request *http.Request) {
	session, err := store.Get(request, "zhengwuyangguang")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(session.Values["name"])
	fmt.Println(session.Values["height"])
}

func handleSetSession(writer http.ResponseWriter, request *http.Request) {
	session, err := store.Get(request, "zhengwuyangguang")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["name"] = "alex"
	session.Values["age"] = 19
	session.Save(request, writer)
}

func handleForm(writer http.ResponseWriter, request *http.Request) {
	executeTemplate := templates.ExecuteTemplate(writer, "form.html", nil)
	if executeTemplate != nil {
		panic(executeTemplate)
	}
}

func handleFormAdd(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	comment := request.PostForm.Get("comment")
	templates.ExecuteTemplate(writer, "form.html", comment)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "这是home处理器")
}

func handleDetail(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "这是详情处理器")
}

func handleIndex(writer http.ResponseWriter, request *http.Request) {
	data := struct {
		Title string
		Items []string
	}{
		Title: "胥诺是是个傻子",
		Items: []string{
			"胥诺是谁",
			"胥诺是傻子",
		},
	}
	templates.ExecuteTemplate(writer, "index.html", data)
}

func handleDb(writer http.ResponseWriter, request *http.Request) {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test_new")
	if err != nil {
		return
	}

	rows, err := db.Query("select id,title from news")
	if err != nil {
		return
	}

	var (
		id    string
		title string
	)

	var news []string

	for rows.Next() {
		rows.Scan(&id, &title)
		news = append(news, title)
	}
	templates.ExecuteTemplate(writer, "db.html", news)
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

var store = sessions.NewCookieStore([]byte("test"))

var db *sql.DB
