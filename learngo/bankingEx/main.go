package main

import (
	"fmt"

	"banking/banking"
)

func main() {
	account := banking.NewAccount("junwoo")
	account.Deposit(1000)
	account.Balance()
	err := account.WithDraw(2000)
	if err != nil {
		//log.Fatalln(err) 프로그램 종료
		fmt.Println(err)
	}
	account.Balance()
	fmt.Println("owner is", account.Owner())
	// struct 의 String() 메소드를 호출하는 것임
	fmt.Println(account)

}
