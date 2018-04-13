package main

import (
	"log"
	"quwb_jobs/dbutil"
	"quwb_jobs/task/shopalloc"
)

func main() {
	db, err := dbutil.GetDB()
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer db.Close()

	//创建任务实例
	task1 := shopalloc.NewAllocShop(db)

	//任务开启
	go task1.Start()

	select {}
}
