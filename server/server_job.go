package main

import (
	"quwb_jobs/common"
	"log"
	"database/sql"
	"time"
	"quwb_jobs/task"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	c, err := common.ReadConf("./config.json")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	//连接字符串
	//connString := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", server, port, database, user, password)

	if db, err := sql.Open("mssql", c.Db["sqlserver"]); err == nil{
		defer db.Close()

		fmt.Println("===========================循环开始===========================")
		t := task.NewAllocShop(db)
		for{
			time.Sleep(5 * time.Minute)
			t.Exec()
		}
		fmt.Println("===========================循环结束===========================")
	}else{
		log.Println(err)
	}
}
