package task

import (
	"database/sql"
	"fmt"
	"github.com/robfig/cron"
)

type AllocShop struct {
	db *sql.DB
}

func NewAllocShop(db *sql.DB) *AllocShop {
	return &AllocShop{
		db: db,
	}
}

func (a *AllocShop) Start() {
	c := cron.New()
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		a.exec()
	})

	c.Start()
}

func (a *AllocShop) exec()  {
	fmt.Println("===========================分配任务开始===========================")
	fmt.Println("===========================分配任务结束===========================")
}
