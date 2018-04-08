package main

import (
	"log"
	"quwb_jobs/dbutil"
	"quwb_jobs/task"
)

func main() {
	db, err := dbutil.GetDB()
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer db.Close()

	//创建任务实例
	shopTask := task.NewAllocShop(db)

	//任务开启
	go shopTask.Start()

	select {}
}
