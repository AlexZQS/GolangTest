package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/first")
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}

	defer func() {
		if db != nil {
			db.Close()
		}
	}()
	db.Ping()
	stmt, err := db.Prepare("select * from people")
	if err != nil {
		fmt.Println("预处理失败")
		return
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("查询失败")
		return
	}

	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	for rows.Next() {
		var id int
		var name string
		var address string
		rows.Scan(&id, &name, &address)
		fmt.Printf("id = %d name = %s address = %s", id, name, address)
	}

}
