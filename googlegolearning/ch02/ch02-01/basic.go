/*
2-1 变量定义
使用var关键字
	定义变量，没有赋值，系统默认赋零值：var a, b, c bool
	定义变量，同时赋初值，并且一条语句可以给多个变量进行定义和赋值：var s1, s2 string = "hello, "world"
	使用var关键字定义变量可以可放在函数内，或直接放在包内
	在包内或者函数内都可以使用var()集中定义变量
让编译器自动决定类型
	var a, b, i, s1, s2 = true,false, 3, "hello", "world"

使用:=定义变量，只能在函数内使用
	a, b, i, s1, s2 := true, false, 3, "hello', "world"

*/
package main

import "fmt"

//go语言没有全局变量的说法，只有包变量
var aa = 3

//通过括号的方式进行包变量定义
var (
	bb = 4
	cc = true
	dd = "hello"
)

func main() {
	varZeroValue()
	varInitialValue()
	varTypeDeduction()
	varShorter()
	fmt.Println(aa, bb, cc, dd)
}

// 在go语言中变量名写在变量类型的前面
// fmt.Printf 带格式打印，%q是Quotation的意思，会把字符串的引号打印出来
func varZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

//定义变量并且赋初始值，而且一条语句可以定义多个变量并赋初始值
func varInitialValue() {
	var a, b int = 3, 4
	fmt.Println(a, b)
}

//go语言可以根据值推断出类型，并且同一行可以赋值不同的类型
func varTypeDeduction() {
	var a, b, c, d = 3, 4, true, "hello"
	fmt.Println(a, b, c, d)
}

//在函数中可以使用短变量定义并赋值的方式
func varShorter() {
	a, b, c, d := 3, 4, true, "hello"
	fmt.Println(a, b, c, d)
}
