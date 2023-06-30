package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(res, res.Body)

	// bs := []byte{}
	bs := make([]byte, 99999)
	res.Body.Read(bs)
	fmt.Println(string(bs))
}