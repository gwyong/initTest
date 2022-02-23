package main

import (
	"fmt"
	"strings"
)

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

func lenAndUpper(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	defer fmt.Println("I am done.")
	return
}
func main() {
	totalLength, up := lenAndUpper("yong")
	fmt.Println(totalLength, up)
}
