package main

import (
	"github.com/Seiya-Tagami/go-practice/iij-bootcamp/7/shop"
	"net/http"
)

func main() {
	myshop := shop.NewGyudon()
	http.HandleFunc("/", myshop.Eat)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}
