package main

import (
	"fmt"
	"sync"
)

var boolChan = make(chan bool)
var otrChan = make(chan string)
var wg sync.WaitGroup

func main() {
	go a()
	go b()

	wg.Add(1)
	otrChan <- "msg"
	wg.Add(1)
	otrChan <- "msg"
	close(boolChan)
	wg.Add(1)
	otrChan <- "msg"

	wg.Wait()
}

func a() {
	fmt.Println("This execustes sometimes if wg.Wait() becomes slow!")
	for {
		_, ok := <-boolChan
		if !ok {
			fmt.Println("routine a closed")
			return
		}
	}
}

func b() {
	for {
		// Note: Must be inside for{} to be executed
		select {
		case _, ok := <-boolChan:
			if !ok {
				fmt.Println("routine b closed")
				return
			}
		default:
		}

		// fmt.Println(">") // Important: If we enable this then we get the expected error. Because select{} is more time consuming that just <- so message is received fester than the close signal and wg.Done() completes and the main go routine dies.
		msg := <-otrChan
		wg.Done()
		fmt.Println(msg)

		// Alt: This is better if possible
		// select { // Use select if needs to listen from multiple chans
		// case msg := <-otrChan:
		// 	wg.Done()
		// 	fmt.Println(msg)
		// case _, ok := <-boolChan:
		// 	if !ok {
		// 		fmt.Println("routine b closed")
		// 		return
		// 	}
		// }

		/* msg := <-otrChan // Deadlock
		fmt.Println(msg)

		_, ok := <-boolChan // Go stuck here waiting for boolChan and cannot receive from otrChan
		if !ok {
			fmt.Println("routine b closed")
			return
		} */
	}
}
