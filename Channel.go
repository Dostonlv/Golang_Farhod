package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan int)
	go getRandomNumber(channel)
	for randomNumber := range channel {
		fmt.Println("tasodifiy son ", randomNumber)
	}
}
func getRandomNumber(channel chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 3; i++ {
		number := rand.Intn(1000)
		time.Sleep(time.Second * 1)
		channel <- number
	}
	close(channel)
}

//func main() {
//	//version 1
//	//channel := make(chan string, 2)
//	//channel <- "anor"
//	//channel <- "olma"
//	//fmt.Println(<-channel)
//	//fmt.Println(<-channel)
//	//version 2
//	channel := make(chan string)
//	go func() {
//		channel <- "anor"
//		channel <- "olma"
//		channel <- "nok"
//	}()
//	fmt.Println(<-channel)
//	fmt.Println(<-channel)
//	fmt.Println(<-channel)
//}
