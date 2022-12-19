package main

import (
	"fmt"
)

func main() {
	testArrays3()
}

func testArrays1() {
	var myarr [3]string

	myarr[0] = "GO"
	myarr[1] = "Virtualdars"
	myarr[2] = "Dasturlash"

	fmt.Println("Qatorning elementlari:")
	fmt.Println("1:", myarr[0])
	fmt.Println("2:", myarr[1])
	fmt.Println("3:", myarr[2])
}
func testArrays2() {
	myarr := [3]int{1, 2, 3}
	fmt.Println(myarr)
}

func testArrays3() {
	myarr1 := [...]int{1, 4, 3, 2}
	myarr2 := [4]int{1, 4, 3, 2}
	fmt.Println(myarr1 == myarr2)
	myslice := myarr2[1:4]
	fmt.Println(myslice)
}
