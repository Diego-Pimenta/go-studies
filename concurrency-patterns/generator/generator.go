package main

import "fmt"

func main() {
	ch := write("Hello, World!")

	for range 10 {
		fmt.Println(<-ch)
	}
}

func write(text string) <-chan string { // encapsula a chamada para uma goroutine e retorna um canal
	ch := make(chan string)

	go func() {
		for {
			ch <- fmt.Sprintf("Valor recebido: %s", text)
		}
	}()

	return ch
}
