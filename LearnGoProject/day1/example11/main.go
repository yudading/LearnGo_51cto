package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var n int

	n = rand.Intn(100)

	for {
		var input int
		fmt.Scanf("%d", &input)
		flag := false
		switch {
		case input == n:
			fmt.Println("you are right")
			flag = true
		case input > n:
			fmt.Println("bigger")
		case input < n:
			fmt.Println("small")
		}
		if flag {
			break
		}
	}
}
