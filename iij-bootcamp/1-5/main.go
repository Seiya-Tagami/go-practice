package main

import (
	"fmt"
	"github.com/Seiya-Tagami/go-basics/iij-bootcamp/1-5/monkey"
	"os"
)

func main() {
	var name1 string = "GYUDON"
	if _, err := monkey.Eat(name1); err != nil {
		fmt.Fprintf(os.Stderr, "cannot eat: %s\n", err)
	}

	var name2 string = ""
	if _, err := monkey.Eat(name2); err != nil {
		fmt.Fprintf(os.Stderr, "cannot eat: %s\n", err)
	}
}
