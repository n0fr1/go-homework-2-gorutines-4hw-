package main

import (
	"fmt"
	"time"
)

func main() {

	var workers = make(chan int, 1)
	var countChan = make(chan int, 1)
	countChan <- 0

	for i := 0; i < 1000; i++ {
		workers <- i

		go func() {
			defer func() {
				<-workers
			}()

			counter := <-countChan
			counter++
			countChan <- counter
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Printf("%d", <-countChan)
}
