package main

import (
	"fmt"
	"list/list"
)

func main() {
	l := list.List{}
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.Add(6)
	l.Add(4)
	l.Print_All()
	fmt.Println("-------------")
	l.Clear()
	l.Print_All()

}

// func sum(a, b int) (sum int) {
// 	sum = a + b
// 	return sum
// }

// type Person struct {
// 	Name string
// 	Age  int
// }
