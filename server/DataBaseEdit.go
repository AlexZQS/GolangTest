package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := sql.Open("mysql", "root:1234@tcp(localhost:3306)/first")
	defer func() {
		if db != nil {
			db.Close()
		}
	}()
	stmt, _ := db.Prepare("update people set name =?,address=? where id = ?")
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	result, _ := stmt.Exec("王五一", "北京", 1)
	count, _ := result.RowsAffected()
	if count > 0 {
		fmt.Println("修改成功")
	} else {
		fmt.Println("修改失败")
	}

}
