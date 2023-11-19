package main

import (
	"fmt"
	"os"

	"github.com/Seiya-Tagami/go-basics/iij-bootcamp/6/shop"
)

func main() {
	myshop := shop.NewGyudon()
	if _, err := myshop.Eat(); err != nil {
		fmt.Fprintf(os.Stderr, "cannot eat: '%s'\n", err)
	}
}
