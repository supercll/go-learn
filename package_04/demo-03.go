package main

import "fmt"

func main() {
	demo02()
}

func demo() {
	var dict map[string]int = map[string]int{}
	dict["a"] = 1
	fmt.Println(dict)

	var dict2 = map[string]int{}
	dict2["b"] = 2
	fmt.Println(dict2)

	dict3 := map[string]int{"test": 0}
	dict3["c"] = 3
	fmt.Println(dict2)

	dict4 := make(map[string]int)
	dict4["d"] = 4
	fmt.Println(dict4)
}

func demo01() {
	// 构造一个map
	str := "aba"
	dict := map[rune]int{}
	for _, value := range str {
		dict[value]++
	}
	fmt.Println(dict) // map[97:2 98:1]

	// 访问map里不存在的key，并不会像C++一样自动往map里插入这个新key
	value, ok := dict['z']
	fmt.Println(value, ok) // 0 false
	fmt.Println(dict)      // map[97:2 98:1]

	// 访问map里已有的key
	value2 := dict['a']
	fmt.Println(value2) // 2
}
func demo02() {
	dict := map[string]int{"a": 1, "b": 2}
	fmt.Println(dict) // map[a:1 b:2]

	// 删除"a"这个key
	delete(dict, "a")
	fmt.Println(dict) // map[b:2]

	// 删除"c"这个不在的key，对map结果无影响
	delete(dict, "c")
	fmt.Println(dict) // map[b:2]
}
