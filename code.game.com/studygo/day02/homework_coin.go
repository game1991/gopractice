package main

import "fmt"

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth
分配规则如下：
a.名字中每包含1个'e'或者'E'分1枚金币
b.名字中每包含1个'i'或者'I'分2枚金币
c.名字中每包含1个'o'或者'O'分3枚金币
d.名字中每包含1个'u'或者'U'分4枚金币
e.名字中每包含1个'z'或者'Z'分5枚金币

写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现'dispathCoin'函数
*/
var (
	coins        = 50
	users        = []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispathCoin()
	fmt.Println("剩余多少金币：", left)
	for k, v := range distribution {
		fmt.Printf("%s -> %d\n", k, v)
	}
}

func dispathCoin() (left int) {

	//1.依次拿到每个人的名字
	for _, name := range users {
		//2.拿到一个人名根据分金币的规则去分金币
		for _, c := range name {
			//2.1每个人的分金币数储存在distribution中
			//2.2记录剩下的金币数
			switch c {
			case 'e', 'E':
				distribution[name]++
				coins--
			case 'i', 'I':
				distribution[name] += 2
				coins -= 2
			case 'o', 'O':
				distribution[name] += 3
				coins -= 3
			case 'u', 'U':
				distribution[name] += 4
				coins -= 4
			case 'z', 'Z':
				distribution[name] += 5
				coins -= 5
			}
		}

	}
	//3.整个2步骤执行完成就可以得出最终每个人分的金币数和剩余金币
	left = coins
	return left
}
