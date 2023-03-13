package main

import (
	"fmt"

	qt "github.com/weimenglee/stringmod/quotes"
	str "github.com/weimenglee/stringmod/strings"
)

func main() {
	o, e := str.CountOddEven("12345678")
	fmt.Println(o, e)

	fmt.Println(qt.GetEmoji("turtle"))
}
