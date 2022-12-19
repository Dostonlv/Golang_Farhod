package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		display1("Salom berdik")
		wg.Done()
	}()
	go func() {
		display1("Alik oldik")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("completed")

}

func display1(input string) {
	for i := 1; i <= 7; i++ {
		fmt.Println(i, input)
		time.Sleep(1 * time.Second)
	}

}
