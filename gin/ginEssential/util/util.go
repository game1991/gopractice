package util

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {

	//给随机数时一定要注意设置seed
	rand.Seed(time.Now().Unix())

	letter := []byte("asdfghjklqwertyuiopzxcvbnmASDFGHJKLQWERTYUIOPZXCVBNM")
	result := make([]byte, n)
	for i := range result {
		result[i] = letter[rand.Intn(len(letter))]
	}
	return string(result)
}
