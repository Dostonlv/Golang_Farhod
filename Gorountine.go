package main

import (
	"fmt"
	"time"
)

func main() {
	go display("Salom berdik")
	go display("Alik oldik")

	fmt.Scanln()
}

func display(input string) {
	for i := 1; true; i++ {
		fmt.Println(i, input)
		time.Sleep(1 * time.Second)
	}
}
