package main

import (
	"fmt"
)

var ca = make(chan int)

func add1() {

	_ :<-ca
	fmt.Println("add1")
}

func add2() {

	_: <-ca
	fmt.Println("add2")
}

func main() {
	fmt.Println("hello")
	var a int =1
	{
		a=2
	}

	fmt.Println(a)

}
