package dbutil

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	_ "odbc/driver"
	"flag"
	"fmt"
	"log"
)

var (
	debug    = flag.Bool("debug", true, "enable debuging")
	server   = flag.String("server", "192.168.0.115", "the database server")
	port     = flag.Int("port", 1433, "the database port")
	user     = flag.String("user", "sa", "the database user")
	password = flag.String("password", "qu90()op", "the database password")
	database = flag.String("database", "ShopHouse", "the database name")
)

//获取sql.DB对象
func GetDB() (*sql.DB, error) {
	if *debug {
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" port:%s\n", *port)
		fmt.Printf(" user:%s\n", *user)
		fmt.Printf(" password:%s\n", *password)
	}

	connString := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%d;encrypt=disable",
		*server, *database, *user, *password, *port)

	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}

	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
		return db, err
	}

	err = db.Ping()

	if err != nil {
		fmt.Print("PING:%s", err)
		return db, err
	}

	return db, nil
}
