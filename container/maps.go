package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "cui",
		"course":  "go",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil

	fmt.Println(m)
	fmt.Println(m2)
	fmt.Println(m3)

	fmt.Println("Traverse map")
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k := range m {
		fmt.Println(k)
	}

	fmt.Println("Get values。")
	courseName := m["course"]
	fmt.Println(courseName)
	errorValue := m["errorValue"] // 不存在的key返回空
	fmt.Println(errorValue)

	if courseName, ok := m["course"]; ok {
		// 判断值是否存在
		fmt.Println(courseName, ok)
	} else {
		fmt.Println("key des not exist")
	}

	fmt.Println("Delete values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")

	name, ok = m["name"]
	fmt.Println(name, ok)

}
