package main

import (
	"net/http"

	"github.com/Seiya-Tagami/go-basics/iij-bootcamp/7/shop"
)

func main() {
	myshop := shop.NewGyudon()
	http.HandleFunc("/", myshop.Eat)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}
