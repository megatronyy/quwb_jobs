package main

import (
	//"quwb_jobs/common"
	//"log"
	//"database/sql"
	//"time"
	//"quwb_jobs/task"
	"fmt"

	//_ "github.com/denisenkom/go-mssqldb"
	"github.com/robfig/cron"
)

func main() {
	isStop := make(chan bool)
	/*c, err := common.ReadConf("./config.json")
	if err != nil {
		log.Println(err)
		panic(err)
	}*/

	i := 0
	c := cron.New()
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		fmt.Println("===========================分配任务开始===========================")
		if i == 5 {
			isStop <- true
		}
		i++
		fmt.Println("===========================分配任务结束===========================")
	})

	c.Start()

	<-isStop
	//连接字符串
	//connString := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", server, port, database, user, password)

	/*if db, err := sql.Open("mssql", c.Db["sqlserver"]); err == nil{
		defer db.Close()

		fmt.Println("===========================循环开始===========================")
		t := task.NewAllocShop(db)
		for{
			t.Exec()
			time.Sleep(5 * time.Minute)
		}
		fmt.Println("===========================循环结束===========================")
	}else{
		log.Println(err)
	}*/
}
