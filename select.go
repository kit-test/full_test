package main

import (

	"time"
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {

		select { //阻塞在select处
		case num1 := <-c1:
			fmt.Println("received", num1)
		case num2 := <-c2:
			fmt.Println("received", num2)

		}

	}
}

