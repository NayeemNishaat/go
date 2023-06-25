package main

import "fmt"

func main() {
	// var colors map[string]string // Note: [type]value

	// colors := make(map[string]string)

	// colors["white"] = "#fff"
	// delete(colors, "white")

	colors := map[string]string{
		"red":   "#f00",
		"green": "#0f0",
		"blue":  "#00f",
	}

	printMap(colors)
	// fmt.Println(colors)
}

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Println(k, v)

	}

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }
}
