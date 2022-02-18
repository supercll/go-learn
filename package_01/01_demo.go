package main

import "fmt"

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	hello_world()
	println()
	println(a, b, c)

}

func hello_world() {
	const STRING string = "hello world"
	println()
	print(STRING)
}

func GetData() (int, int) {
	return 100, 200
}
