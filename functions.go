package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	var x, y int
	fmt.Scanln(&x, &y)
	//res := power(2, 3)
	//fmt.Println(res)
	result := power(x, y)
	fmt.Println(result)

}

func sum(x int, y int) int {
	return x + y
}
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Manfiy son uchun aniq emas")
	}
	return math.Sqrt(x), nil
}
func power(x, y int) int {

	num := 1
	for i := 1; i <= y; i++ {
		num = num * x
	}
	return num
}
