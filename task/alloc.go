package task

import (
	"database/sql"
	"fmt"
	"github.com/robfig/cron"
)

type AllocShop struct {
	db *sql.DB
}

func (a *AllocShop) Start() error {
	i := 0
	c := cron.New()
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		a.Exec()
	})

	c.Start()
}

func NewAllocShop(db *sql.DB) *AllocShop {
	return &AllocShop{
		db: db,
	}
}

func (a *AllocShop) Exec() error {
	var err error
	fmt.Println("===========================分配任务开始===========================")
	fmt.Println("===========================分配任务结束===========================")

	return err
}
