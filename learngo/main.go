package main

import (
	"fmt"

	"github.com/gwyong/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "word 1"
	definition := "definition 1"
	dictionary.Add(word, definition)
	dictionary.Search(word)
	dictionary.Delete(word)
	def, err := dictionary.Search(word)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(def)

}
