package main

import "fmt"

func main() {
	var a int = 0

	switch a {

	case 0:
		fmt.Println("a is equals 0")
		fmt.Println("yes")
	case 10:
		fmt.Println("a is equals 10")
	default:
		fmt.Println("a is equals default")
	}

	g := grade(101)
	fmt.Println(g)
}

func grade(score int) string {
	g := ""

	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprint("wrong score:%d", score))
	case score < 60:
		g = "f"
	case score < 80:
		g = "c"
	case score < 90:
		g = "b"
	case score <= 100:
		g = "a"
	}

	return g
}
