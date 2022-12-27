package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Init() {

}
/*
Get Led Status
@param : db
@return : *LedStatus
*/
func GetLedStatus ( db *sql.DB )  (*LedStatus,error) {
	row := db.QueryRow("select * from led_status where name = 'led' ")
	var status LedStatus
	err := row.Scan(&status.Status,&status.Controller,&status.Name,&status.Id,&status.Updated)
	
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &status,nil
}
/*
Update Led Status 
@param : db, order
@return : rowAffected , err
*/
func UpdateLedStatus (db *sql.DB, order string ) (int64,error) {
	
	if !(order == "on" || order == "off") {
		return 0,errors.New("CAN NOT UPDATE LED STATUS")
	}
	fmt.Println(order)
	result,err := db.Exec("update led_status set status = ? where name = 'led'", order)
	if err != nil {
		log.Fatal(err)
	}
	rowAffected,err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	
	return rowAffected , nil
}

func ExampleCode () {
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
	status,err := GetLedStatus(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(status)
	result,err := UpdateLedStatus(db,"on")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	defer db.Close()
}