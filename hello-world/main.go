package main

import "fmt"

func init() {
	fmt.Println("INIT First")
}

func main() {
	fmt.Println("Hello Last")
}

func init() {
	fmt.Println("INIT Second")
}
