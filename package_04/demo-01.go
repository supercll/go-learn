package main

import (
	"fmt"
	"reflect"
)

func printSlice(param []int) {
	fmt.Printf("slice len:%d, cap:%d, value:%v\n", len(param), cap(param), param)
}

func main() {
	// demo01()
	// demo02()
	// demo03()
	// demo()
	demo04()
}

func demo() {
	slice1 := []int{1}
	slice2 := make([]int, 3, 100)
	printSlice(slice1)
	printSlice(slice2)
}

func demo01() {

	slice3 := []int{}
	var slice4 []int
	fmt.Println("is nil: ", slice3 == nil, slice4 == nil)
}

func demo02() {
	slice := make([]int, 3, 10)
	/*下标访问切片*/
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d]=%d\n", i, slice[i])
	}

	/*range迭代访问切片*/
	for index, value := range slice {
		fmt.Printf("slice[%d]=%d\n", index, value)
	}
}

func demo03() {
	// 对数组做切片
	array := [3]int{1, 2, 3} // array是数组
	slice3 := array[1:3]     // slice3是切片
	fmt.Println("slice3 type:", reflect.TypeOf(slice3))
	fmt.Println("slice3=", slice3) // slice3= [2 3]

	slice4 := slice3[1:2]
	fmt.Println("slice4=", slice4) // slice4= [3] 左开右闭区间

	/* slice5->slice4->slice3->array
	切片后的指针会指向原切片的数组空间，会影响原数据
	对slice5的修改，会影响到slice4, slice3和array
	*/
	slice5 := slice4[:]
	fmt.Println("slice5=", slice5) // slice5= [3]

	slice5[0] = 10
	fmt.Println("array=", array)   // array= [1 2 10]
	fmt.Println("slice3=", slice3) // slice3= [2 10]
	fmt.Println("slice4=", slice4) // slice4= [10]
	fmt.Println("slice5=", slice5) // slice5= [10]
}

func change1(param []int) {
	param[0] = 100             // 这个会改变外部切片的值
	param = append(param, 200) // append不会改变外部切片的值
}

func change2(param *[]int) {
	*param = append(*param, 300) // 传切片指针，通过这种方式append可以改变外部切片的值
}

func demo04() {
	slice := make([]int, 2, 100)
	fmt.Println(slice) // [0, 0]

	change1(slice)
	fmt.Println(slice) // [100, 0]

	change2(&slice)
	fmt.Println(slice) // [100, 0, 300]
}
