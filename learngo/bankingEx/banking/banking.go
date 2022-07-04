package banking

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("can't withDraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit 예금 , Pointer Reciever ,struct 의 pointer 를 가져온다.
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// WithDraw 계좌 인출
func (a *Account) WithDraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil

}

// Balance 잔액 , Reciver ,struct 의 복사본을 가져온다.
func (a Account) Balance() int {
	fmt.Println(a.balance)
	return a.balance
}

func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}
func (a Account) Owner() string {
	return a.owner
}

// struct print 시 호출되는 메소드
func (a Account) String() string {
	return fmt.Sprint(a.owner, "'s account \n has", a.balance)
}
