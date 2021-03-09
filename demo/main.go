package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	const url = "http://www.zhenai.com/zhenghun"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("status code error:", response.StatusCode)
		return
	}

	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s/n", all)

}
