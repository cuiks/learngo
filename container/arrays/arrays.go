package main

import "fmt"

func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]bool

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	sum := 0
	for _, v := range arr3 {
		sum += v
	}
	fmt.Println(sum)

	printArray(arr3)
	printArray(arr1)
	// 数组是值传递。当做参数传递会拷贝数组
	fmt.Println(arr1, arr3)

	// go语言一般不直接使用数组。切片

}
