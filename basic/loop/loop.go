package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

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
	printFile("basic/branch/abc.txt")
	s := `asd""asd
	sd\n
	
	asdasd`
	printFileContents(strings.NewReader(s))

	//forever()
}
