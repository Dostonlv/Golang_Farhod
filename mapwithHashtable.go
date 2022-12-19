package main

import "fmt"

func main() {
	statuses := make(map[string]int)

	statuses["pending"] = 0
	statuses["initied"] = 1
	statuses["running"] = 2
	statuses["timeout"] = 3
	statuses["successful"] = 4
	statuses["failed"] = 5
	var successfulstatus = statuses["successful"]
	fmt.Println(successfulstatus)
	fmt.Println(statuses)
	delete(statuses, "failed")
	fmt.Println(statuses)

}
