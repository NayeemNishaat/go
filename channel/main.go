package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) // Remark: Creating new go routines when the child routine is blocked.
	}

	// fmt.Println(<-c) // Important: Blocking call, waiting for messages from the channel
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	for {
		go checkLink(<-c, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // Note: Blocking call
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link

		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
