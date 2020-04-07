package main

import (
	"fmt"
	"log"
)

var printFN = func(idx int, bean interface{}) error {
	fmt.Printf("%d:%#v\n", idx, bean.(*Account))
	return nil
}

func main() {
	fmt.Println("Welcome to bank of xorm")
	count, err := getAccount()
	if err != nil {
		log.Fatal("Fail to get Account count:", err)
		return
	}
	fmt.Println("Account count:", count)

	for i := count; i < 10; i++ {
		if err := newAccount(fmt.Sprintf("joe%d", i), float64(i*100.00)); err != nil {
			log.Fatal("Fail to create Account,err:", err)
			return
		}
	}

	fmt.Println("Query all records:")
	x.Iterate(new(Account), printFN)

}
