package main

import "fmt"

// 定义动物类接口
type Animal interface {
	speak()
}

// cat结构体
type Cat struct {
	name string
	age  int
}

// cat方法
func (cat Cat) speak() {
	fmt.Println("cat miaomiaomiao")
}

// dog结构体
type Dog struct {
	name string
	age  int
}

// Dog方法
func (dog *Dog) speak() {
	fmt.Println("dog wangwangwang")
}

func main() {
	var animal Animal = Cat{"gaffe", 1}
	animal.speak() // cat miaomiaomiao

	/*
	   如果用了指针接受者，那给interface变量赋值的时候要传指针
	*/
	animal = &Dog{"caiquan", 2}
	animal.speak() // dog wangwangwang
}
