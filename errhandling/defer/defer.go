package main

import (
	"bufio"
	"fmt"
	"learngo/functional/fib"
	"os"
)

func tryDefer() {
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(3)
	////return
	//panic("error")
	//fmt.Println(4)

	for i := 0; i < 100; i++ {
		defer fmt.Println(i)

		if i == 30 {
			panic("error")
		}
	}
}

//func writeFile(fileName string) {
//	file, err := os.Create(fileName)
//	if err != nil {
//		panic(err)
//	}
//	defer file.Close()
//
//	writer := bufio.NewWriter(file)
//	defer writer.Flush()
//
//	f := fib.Fibonacci()
//	for i := 0; i < 20; i++ {
//		fmt.Fprintln(writer, f())
//	}
//}

func writeFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op)
			fmt.Println(pathError.Path)
			fmt.Println(pathError.Err)
		}
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	fmt.Println("###")
	writeFile("fib.txt")
	fmt.Println("###")
	//tryDefer()
}
