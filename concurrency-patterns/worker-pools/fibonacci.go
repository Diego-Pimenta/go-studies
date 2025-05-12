package main

func main() {
	batch := make(chan int, 45)
	results := make(chan int, 45)

	go worker(batch, results)
	go worker(batch, results)

	for i := range 45 {
		batch <- i // ao popular batch, o worker vai processar os dados
	}
	close(batch)

	for range 45 {
		result := <-results
		println(result)
	}
}

func worker(batch <-chan int, results chan<- int) { // batch so traz dados e results armazena os dados
	for num := range batch {
		results <- fibonacci(num)
	}
}

func fibonacci(num int) int {
	if num <= 1 {
		return num
	}
	return fibonacci(num-1) + fibonacci(num-2)
}
