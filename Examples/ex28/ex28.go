package main

import (
	"fmt"
	"strings"
)

func reformat(message string, formatter func(string) string) string {
	format1 := formatter(message)
	format2 := formatter(format1)
	format3 := formatter(format2)
	
	return "TEXTIO: " + format3
}


func main() {
	result := reformat("hello world", strings.ToUpper)
	fmt.Println(result)
}