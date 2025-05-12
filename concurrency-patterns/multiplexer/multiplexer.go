package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := multiplexer(write("Hello World!"), write("Hi there"))

	for range 10 {
		fmt.Println(<-ch)
	}

}

func multiplexer(ch1, ch2 <-chan string) <-chan string { // simplifica 2 canais em 1
	ch := make(chan string)

	go func() {
		for {
			select {
			case msg1 := <-ch1:
				ch <- msg1
			case msg2 := <-ch2:
				ch <- msg2
			}
		}
	}()

	return ch
}

func write(text string) <-chan string {
	ch := make(chan string)

	go func() {
		for {
			ch <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return ch
}
