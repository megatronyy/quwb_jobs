package main

import (
	"quwb_jobs/config"
	"log"
	"database/sql"
	"time"
	"quwb_jobs/server/task"
	"fmt"
)

func main() {
	c, err := config.ReadConf("./config.json")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	if db, err := sql.Open("adodb", c.Db["sqlserver"]); err == nil{
		defer db.Close()

		fmt.Println("===========================循环开始===========================")
		for{

			time.Sleep(10 * time.Second)

			t := task.NewAllocShop(db)

			go t.Exec()
		}
		fmt.Println("===========================循环结束===========================")
	}else{
		log.Println(err)
	}
}
