package main

import "fmt"

func main() {
	//myarr := [3]string{"yer", "quyosh", "oy"}
	mymap := map[int]string{
		1: "c#",
		2: "golang",
		3: "nodejs"}
	for key, value := range mymap {
		fmt.Println("key: ", key, "value: ", value)
	}
}
