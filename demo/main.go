package main

import "fmt"

func cal(n int) int {
	k := 0
	for n >= 3 {
		k++
		n -= 2
		fmt.Println(n, k)
	}
	return k
}

func main() {
	fmt.Println(cal(10))
}