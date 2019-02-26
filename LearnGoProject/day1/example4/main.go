package main

import "fmt"

func varibleZeroValue() {
	// 变量的声明格式：var <变量名称><变量类型>
	var a int
	var b string
	var c bool

	// 变量的赋值格式：<变量名称>=<表达式>
	a = 10
	c = true
	fmt.Printf("%d,%q,%v", a, b, c)
	fmt.Println()

}

func varibleInitValue() {
	var a, b int = 2, 3
	var s string = "db"
	fmt.Println(a, b, s)
}

func variableInitTypeValue() {
	var a, b, s, c = 2, 3, "db", true

	fmt.Println(a, b, s, c)
}

func variableShorter() {
	a, b, s, c := 2, 3, "db", true
	s = "dddd"
	fmt.Println(a, b, s, c)
}
func main() {

	varibleZeroValue()
	varibleInitValue()
	variableInitTypeValue()
	variableShorter()
}
