package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"strings"
)

func main() {

	b := make([]byte, 10)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// The slice should now contain random bytes instead of only zeroes.
	fmt.Println(bytes.Equal(b, make([]byte, 10)))
	fmt.Println(string(b))

	split()

}

func split() {
	phone := strings.Split("saw123adad_15527966819", "_")[1]
	fmt.Println(phone)
}
