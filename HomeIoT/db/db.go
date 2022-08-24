package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Init() {

	fmt.Println("db")
	db, err := sql.Open("mysql", "junwoo:junwoo123@tcp(junwoodb.clcwfeh6dtye.ap-northeast-2.rds.amazonaws.com:3306)/iotdb")
	if err != nil {
		log.Fatal(err)
	}
	rows,err := db.Query("select * from led_status")
	if err != nil {
		log.Fatal(err)
	}
	columns,err :=rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	for _,col := range columns {

		fmt.Println(col)
	}
	defer db.Close()

}

func GetLedStatus ( db *sql.DB ) {

	db.Query("select * from led_status")
}