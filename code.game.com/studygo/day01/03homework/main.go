package main

import "fmt"

func main() {
	//回文判断
	//上海自来水来自海上 s[0] s[len(s)-1]
	//山西运煤车煤运西山
	//黄山落叶松叶落山黄
	s := "山西运煤车煤运西山"
	//把字符串中的字符放进一个rune切片中
	r := make([]rune, 0, len(s))
	for _, c := range s {
		r = append(r, c)
	}
	fmt.Println("rune切片是：", r)
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")

}
