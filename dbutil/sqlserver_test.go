package dbutil

import (
	"testing"
	"log"
	"fmt"
)

func TestGetDB(t *testing.T) {
	db, err := GetDB()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT TOP 10 * FROM dbo.ACC_SeatGroup")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()

	//通过Statement执行查询
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}

	//建立一个列数组
	cols, err := rows.Columns()
	var colsdata = make([]interface{}, len(cols))

	for i := 0; i < len(cols); i++ {
		colsdata[i] = new(interface{})
		fmt.Print(cols[i])
		fmt.Print("\t")
	}
	fmt.Println()

	for rows.Next() {
		rows.Scan(colsdata...)
	}

	defer rows.Close()
}
