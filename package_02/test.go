package main

import "fmt"

func GetData() (int, int) {
	return 100, 200
}
func main() {
	a, _ := GetData()
	_, b := GetData()
	// b := GetData() // 只用一个变量无法达到想要的效果
	fmt.Println(a, b)
}
