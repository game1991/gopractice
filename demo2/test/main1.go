package main

import (
	"fmt"
	"strings"
)

type Shop struct {
	Name string
}

func main() {
	shops := make([]Shop, 0)
	list := "商品1,商品2,商品3"
	strs := strings.Split(list, ",")
	fmt.Println(strs)
	for _, v := range strs {
		tmp := Shop{Name: v}
		shops = append(shops, tmp)
	}
	fmt.Println(shops)

}
