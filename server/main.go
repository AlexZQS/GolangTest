package main

import (
	"database/sql"
	"fmt"
	//此处不要忘记导入渠道
	_ "github.com/go-sql-driver/mysql"
)

/**
 * @description
 * @time 2019/5/12 22:58
 * @version
 */
func main() {
	//1. 打开链接
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/first")
	db.Ping()

	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}

	//2.预处理SQL
	stmt, err := db.Prepare("insert into people values(default ,?,?)")
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	if err != nil {
		fmt.Println("预处理失败")
		return
	}

	//参数和占位符对应
	result, err := stmt.Exec("张三", "海淀")
	if err != nil {
		fmt.Println("sql执行失败")
		return
	}

	//3.获取结果
	count, err := result.RowsAffected()
	if err != nil {
		fmt.Println("结果获取失败")
		return
	}

	if count > 0 {
		fmt.Println("新增成功")
	} else {
		fmt.Println("新增失败")
	}

	//可能需要获取到新增时主键的值
	id, err := result.LastInsertId()

}
