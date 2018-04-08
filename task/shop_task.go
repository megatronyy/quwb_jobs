package task

import (
	"database/sql"
	"github.com/robfig/cron"
	"log"
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
	spec := "0 */5 * * * ?"
	c.AddFunc(spec, func() {
		a.exec()
	})
	c.Start()
}

func (a *AllocShop) exec() {
	a.queryCustom()
}

func (a *AllocShop) queryCustom() {
	strSql := "SELECT ID,UserID,BusinessID,LocationID,MinPrice,MaxPrice,MinArea,MaxArea FROM [dbo].[Customization] WITH(NOLOCK) WHERE IsActive=1"
	rows, err := a.db.Query(strSql)
	if err != nil {
		log.Fatal("query fail:", err.Error())
	}
	defer rows.Close()

	handleCustom(rows)
}

func handleCustom(r *sql.Rows) {
	var id, userid, businessid, locationid, minarea, maxarea int
	var minprice, maxprice float32

	for r.Next() {
		r.Scan(&id, &userid, &businessid, &locationid, &minprice, &maxprice, &minarea, &maxarea)

		if id > 0 {

		}
	}
}
