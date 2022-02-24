package main

import (
	"fmt"

	"github.com/gwyong/learngo/banking"
)

func main() {
	account := banking.Account{
		Owner:   "Yong",
		Balance: 1000}
	fmt.Println(account)

}
