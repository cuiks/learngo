package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is kshcui@ccxi.com.cn
email2 is abc@qwe.com
email3 is kkk@iiii.com.cn
`

func main() {
	//re, err := regexp.Compile("kshcui@ccxi.com.cn")
	//if err != nil {
	//	panic(err)
	//}
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindString(text)
	fmt.Println(match)
	fmt.Println("####")

	allMatch := re.FindAllString(text, -1)
	fmt.Println(allMatch)
	fmt.Println("####")

	subMatch := re.FindAllStringSubmatch(text, -1)
	for _, m := range subMatch {
		fmt.Println(m)
	}
	fmt.Println("####")

}
