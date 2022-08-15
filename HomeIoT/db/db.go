package db

import (
	"database/sql"
	"fmt"
	"log"
)

func Init() {

	db, err := sql.Open("mysql", "junwoo:junwoo123@tcp(junwoodb.clcwfeh6dtye.ap-northeast-2.rds.amazonaws.com:3306)/db_assignment")
	if err != nil {
		log.Fatal(err)
	}
	rows,err := db.Query("select * from orders")
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
