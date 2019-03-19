package main

import "fmt"

type add_func func(int, int) int

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

//op 指可为任意函数名字 如改为 aa 或 bb

func operator(op add_func, a int, b int) int {
	return op(a, b)
}

//func operator(aa add_func, a int, b int) int {
//	return aa(a,b)
//}

func main() {
	c := add
	d := sub

	sum := operator(c, 5, 6)
	fmt.Println(sum)
	sum1 := operator(d, 5, 6)
	fmt.Println(sum1)

}
