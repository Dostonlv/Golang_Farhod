package main

import (
	"fmt"
	"time"
)

func main() {
	chanel1 := make(chan string)
	chanel2 := make(chan string)

	go func() {
		for {
			chanel1 <- "Tez"
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for {
			chanel2 <- "Sekin"
			time.Sleep(time.Second * 2)
		}
	}()
	for {
		select {
		case message1 := <-chanel1:
			fmt.Println(message1)
		case message2 := <-chanel2:
			fmt.Println(message2)
		default:
			fmt.Println("Malumot yoq")
		}
	}
}
