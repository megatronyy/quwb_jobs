package task

import (
	"database/sql"
	"github.com/robfig/cron"
	"log"
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

//循环订阅信息
func handleCustom(r *sql.Rows) {
	var id, userid, businessid, locationid, minarea, maxarea, minprice, maxprice int32

	for r.Next() {
		r.Scan(&id, &userid, &businessid, &locationid, &minprice, &maxprice, &minarea, &maxarea)

		if id > 0 {
			fmt.Printf("id:%s\t\nuserid:%s\t\nbusinessid:%s\t\n", id, userid, businessid)

			_, err := customProcesser(&Customization{
				ID:         id,
				UserID:     userid,
				BusinessID: businessid,
				LocationID: locationid,
				MinPrice:   minprice,
				MaxPrice:   maxprice,
				MinArea:    minarea,
				MaxArea:    maxarea,
			})

			if err != nil {
				fmt.Printf("customProcesser error:%s\t\n", err.Error())
			}
		}
	}
}

func customProcesser(c *Customization) (bool, error) {
	var ret  bool

	strSql := "SELECT  t1.ShopID" +
	"FROM    dbo.ShopInfo t1 ( NOLOCK )" +
		"LEFT JOIN dbo.UserShopRela t2 ( NOLOCK ) ON t2.ShopID = t1.ShopID" +
	"AND t2.UserID = @UserID" +
	"AND t2.IsActive = 1" +
	"WHERE   t1.LocationID = @LocationID" +
	"AND t1.ShopPrice BETWEEN @MinPrice AND @MaxPrice" +
	"AND t1.ShopArea BETWEEN @MinArea AND @MaxArea" +
	"AND t1.BusinessID = @BusinessID" +
	"AND t2.RelaID IS NULL; "
	fmt.Println(strSql)
	return ret, nil
}
