package main

import (
	"fmt"
)

func main() {
	// demo01()
	// demo02()
	// demo03()
	// demo04()
	demo05()
}

func demo01() {
	var cat int = 1
	var str string = "banana"
	fmt.Printf("%p %p", &cat, &str)
}

func demo02() {
	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"
	// 对字符串取地址, ptr类型为*string
	ptr := &house
	// 打印ptr的类型
	fmt.Println()
	// 打印ptr的指针地址
	fmt.Printf("ptr type: %T\n", ptr) // ptr type: *string
	// 对指针进行取值操作
	fmt.Printf("address: %p\n", ptr) // address: 0xc00008a220
	// 取值后的类型
	value := *ptr
	// 指针取值后就是指向变量的值
	fmt.Printf("value type: %T\n", value) // value type: string
	fmt.Printf("value: %s\n", value)      // value: Malibu Point 10880, 90265
}

// 交换函数
func swap(a, b *int) {
	// 取a指针的值, 赋给临时变量t
	t := *a
	// 取b指针的值, 赋给a指针指向的变量
	*a = *b
	// 将a指针的值赋给b指针指向的变量
	*b = t
}

func demo03() {
	// 准备两个变量, 赋值1和2
	x, y := 1, 2
	z := x
	z2 := *&x
	// 交换变量值
	swap(&x, &y)
	// 输出变量值
	fmt.Println(x, y, "x值赋值给z：", z, "x的指针值指向的变量赋给z2: ", z2)
}

func swap2(a, b *int) {
	a, b = b, a
}

func demo04() {
	x, y := 1, 2
	swap2(&x, &y)
	fmt.Println(x, y)
}

func demo05() {
	str := new(string)
	var ptr *int
	fmt.Println(str, *str, ptr)
	fmt.Println(str == nil, ptr == nil)
	*str = "hello go"
	fmt.Println(str, *str)
}
