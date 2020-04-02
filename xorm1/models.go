package main

import (
	"errors"
	"log"

	"github.com/go-xorm/xorm"
	_"github.com/mattn/go-sqlite3"
)

//Account ...
type Account struct {
	ID      int
	Name    string `xorm:"unique"`
	Balance float64
	Version int `xorm:"version"` //乐观锁
}

var x *xorm.Engine

func init() {
	var err error
	x, err = xorm.NewEngine("sqlite3", "./bank.db")
	if err != nil {
		log.Fatal("Fail to create engine,err:", err)
		return
	}
	if err := x.Sync2(new(Account)); err != nil {
		log.Fatalf("Fail to sync database,err:%#v\n", err)
		return
	}
}

func newAccount(name string, balance float64) error {
	_, err := x.Insert(&Account{Name: name, Balance: balance})
	return err
}

func getAccount(id int) (*Account, error) {
	a := &Account{}
	has, err := x.ID(id).Get(a)
	if err != nil {
		log.Fatal("Fail to get Account,err:", err)
		return nil, err
	} else if !has {
		return nil, errors.New("Not exsist Account")
	}
	return a, nil
}

func makeDeposit(id int, deposit float64) (*Account, error) {
	a, err := getAccount(id)
	if err != nil {
		log.Fatal("Fail to get Account,err:", err)
		return nil, err
	}
	a.Balance += deposit
	if _, err := x.Update(a); err != nil {
		log.Fatal("Update failed,err:", err)
		return nil, err
	}
	return a, nil
}

func makeWithdraw(id int, withdraw float64) (*Account, error) {
	a, err := getAccount(id)
	if err != nil {
		log.Fatal("Fail to get Account,err:", err)
		return nil, err
	}
	if a.Balance <= withdraw {
		return nil, errors.New("No enough withdraw")
	}
	a.Balance -= withdraw
	if _, err := x.Update(a); err != nil {
		log.Fatal("Update failed,err:", err)
		return nil, err
	}
	return a, nil
}

func makeTransfer(id1, id2 int, balance float64) error {
	a1, err := getAccount(id1)
	if err != nil {
		log.Fatal("Fail to get Account,err:", err)
		return err
	}
	a2, err := getAccount(id2)
	if err != nil {
		log.Fatal("Fail to get Account,err:", err)
		return err
	}
	if a1.Balance <= balance {
		return errors.New("No enough withdraw")
	}
	//开启事务
	sess := x.NewSession()
	sess.Begin()
	a1.Balance += balance
	if _, err := x.Update(a1); err != nil {
		sess.Rollback()
		log.Fatal("Update failed,err:", err)
		return err
	}
	a2.Balance -= balance
	if _, err := x.Update(a2); err != nil {
		sess.Rollback()
		log.Fatal("Update failed,err:", err)
		return err
	}
	return sess.Commit()
}

func getAccountAscByid() ([]*Account, error) {
	as := make([]*Account, 0)
	err := x.Desc("id").Find(as)
	return as, err
}

func getAccountAscByBalance() ([]*Account, error) {
	as := make([]*Account, 0)
	err := x.Desc("balance").Find(as)
	return as, err
}

func deleteAccount(id int) error {
	_, err := x.Delete(id)
	if err != nil {
		log.Fatalf("delete %d failed,err:%#v\n", id, err)
		return err
	}
	return nil
}
