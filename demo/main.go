package main

import {
	"fmt"
    "time"
}

func main() {
	messages := make(chan string)

	go func() {
		for i := 1; i <= 5; i++ {
			message := fmt.Sprintf("Message %d", i)
			messages <- message
			fmt.Println("Produced:", message)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			message := <-messages
            fmt.Println("Consumed:", message)
            time.Sleep(time.Second)
		} ()

		time.Sleep(10 * time.Second)
	}
}