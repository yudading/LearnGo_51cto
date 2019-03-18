package main

import "fmt"

func main() {

	Print(10)

}

func Print(n int) {
	for i := 0; i < n+1; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("A")
		}
		fmt.Println()
	}
}
