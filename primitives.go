package main

import (
	"fmt"
	"time"
)

func someFunc(num int){

	fmt.Println(num)

}

func main() {
	go someFunc(1)
	go someFunc(2)
	go someFunc(3)
	time.Sleep(time.Second * 2)
	fmt.Println("Hello")
}
