package task

import (
	"database/sql"
	"fmt"
)

type AllocShop struct {
	db *sql.DB
}

func NewAllocShop(db *sql.DB) *AllocShop {
	return &AllocShop{
		db: db,
	}
}

func (a *AllocShop) Exec() error  {
	var err error
	fmt.Println("===========================分配任务开始===========================")
	fmt.Println("===========================分配任务结束===========================")
	return err
}
