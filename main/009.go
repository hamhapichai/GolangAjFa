package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter amount of line of # you want : ")
	fmt.Scan(&num)
	for a := 0; a < num; a++ {
		fmt.Println("#")
	}
}
