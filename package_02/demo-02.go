package main

import (
	"fmt"
	"math"
)

//全局变量 a
var a int = 13

func main() {
	var f float32 = 16777216 // 1 << 24
	const e = .71828         // 0.71828
	const g = 1.             // 1
	fmt.Println(f == f+1)    // "true"!
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)
	const str = `
	第一行
	第二行
	第三行
	\r\n
	`
	fmt.Println(str)
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point

	// 将常量保存为float32类型
	var c float32 = math.Pi
	// 转换为int类型, 浮点发生精度丢失
	fmt.Println(int(c), c)
}
