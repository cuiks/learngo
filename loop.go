package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertRoBin(n int) string {
	if n <= 0 {
		return strconv.Itoa(n)
	}
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertRoBin(5),  // 101
		convertRoBin(13), // 1101
		convertRoBin(0),
	)
	printFile("abc.txt")
	forever()
}
