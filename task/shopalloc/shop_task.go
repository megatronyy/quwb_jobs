package shopalloc

import (
	"database/sql"
	"github.com/robfig/cron"
	"log"
	"fmt"
	"strconv"
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
	fmt.Println("任务开始：商机分配")
	a.queryCustom()
	fmt.Println("任务结束：商机分配")
}

func (a *AllocShop) queryCustom() {
	strSql := "SELECT ID,UserID,BusinessID,LocationID,MinPrice,MaxPrice,MinArea,MaxArea FROM [dbo].[Customization] WITH(NOLOCK) WHERE IsActive=1"
	rows, err := a.db.Query(strSql)
	if err != nil {
		log.Fatal("query fail:", err.Error())
	}
	defer rows.Close()

	a.handleCustom(rows)

}

//循环订阅信息
func (a *AllocShop) handleCustom(r *sql.Rows) {
	var id, userid, businessid, locationid, minarea, maxarea, minprice, maxprice int

	for r.Next() {
		r.Scan(&id, &userid, &businessid, &locationid, &minprice, &maxprice, &minarea, &maxarea)

		if id > 0 {
			fmt.Printf("用户订阅参数：id:%s\t\nuserid:%s\t\nbusinessid:%s\t\n", id, userid, businessid)

			_, err := a.customProcesser(&Customization{
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

//找到合适的商铺，分发给对应的客户
func (a *AllocShop) customProcesser(c *Customization) (bool, error) {
	strSql := getCustomSQL(c)
	rows, err := a.db.Query(strSql)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	var (
		shopId    int
		shopTitle string
	)
	for rows.Next() {
		rows.Scan(&shopId, &shopTitle)

		if shopId > 0 {
			relaSql := getAddShopRelaSQL(c.UserID, shopId)
			a.db.Exec(relaSql)
			msg := "恭喜获得新商机：" + shopTitle + "，请点击查看详情"
			annSql := getAnnounceSQL(c.UserID, msg)
			a.db.Exec(annSql)
		}
	}

	return true, nil
}

func getCustomSQL(c *Customization) string {
	strSql := "SELECT  t1.ShopID,t1.ShopTilte" +
		" FROM    dbo.ShopInfo t1 ( NOLOCK )" +
		" LEFT JOIN dbo.UserShopRela t2 ( NOLOCK ) ON t2.ShopID = t1.ShopID" +
		" AND t2.UserID = " + strconv.Itoa(c.UserID) +
		" AND t2.IsActive = 1" +
		" WHERE   t1.LocationID = " + strconv.Itoa(c.LocationID) +
		" AND t1.ShopPrice BETWEEN " + strconv.Itoa(c.MinPrice) + " AND " + strconv.Itoa(c.MaxPrice) +
		" AND t1.ShopArea BETWEEN " + strconv.Itoa(c.MinArea) + " AND " + strconv.Itoa(c.MaxArea) +
		" AND t1.BusinessID = " + strconv.Itoa(c.BusinessID) +
		" AND t2.RelaID IS NULL; "
	return strSql
}

func getAddShopRelaSQL(userId, shopId int) string {
	strSql := "INSERT INTO dbo.UserShopRela(UserID,ShopID,CreateTime,IsActive)" +
		" VALUES(" + strconv.Itoa(userId) + " ," + strconv.Itoa(shopId) + " , GETDATE() ,1 )"

	return strSql
}

func getAnnounceSQL(userId int, context string) string {
	strSql := "INSERT INTO dbo.Announce(UserID,Context)" +
		" VALUES(" + strconv.Itoa(userId) + " ,'" + context + "' )"

	return strSql
}
