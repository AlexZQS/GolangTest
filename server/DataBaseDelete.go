package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/first")
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}

	db.Ping()
	stmt, err := db.Prepare("delete from people where id = ?")
	if err != nil {
		fmt.Println("预处理失败")
		return
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	result, err := stmt.Exec(2)
	if err != nil {
		fmt.Println("执行失败")
		return
	}
	count, _ := result.RowsAffected()
	if count > 0 {
		fmt.Println("删除成功")
	} else {
		fmt.Println("删除失败")
	}
}
