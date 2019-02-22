package aaa

import "fmt"

func Xxx(n int) {
	for i := 0; i <= n; i++ {
		fmt.Printf("%d + %d = %d\n", i, n-i, n)
	}
}
