package main

import (
	"fmt"
	"time"
)

func printNumber1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func printLetters1() {
	for i := 'A'; i < 'A'+10; i++ {
		fmt.Printf("%c ", i)
	}
}

func printNumbers2(w chan bool) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	w <- true
}

func printLetters2(w chan bool) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	w <- true
}

func print1() {
	printNumber1()
	printLetters1()
}

func goPrint1() {
	go printNumber1()
	go printLetters1()
}

func thrower(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println("throw >>", i)
	}
}

func cather(c chan int) {
	for i := 0; i < 5; i++ {
		num := <-c
		fmt.Println("catch <<", num)
	}
}

func main() {
	c := make(chan int, 3)
	go thrower(c)
	go cather(c)
	time.Sleep(100 * time.Millisecond)
}
