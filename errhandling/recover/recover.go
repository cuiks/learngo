package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("error occurred: ", err)
		} else {
			panic(fmt.Sprintf("dont know: %s", r))
		}
	}()

	//panic(errors.New("this is an error"))
	//b := 0
	//a := 1 / b
	//fmt.Println(a)
	panic(123)
}

func main() {
	tryRecover()
}
