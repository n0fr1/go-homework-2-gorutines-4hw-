package main

import (
	"fmt"
	"time"
)

func main() {

	var workers = make(chan int, 100)

	for i := 0; i <= 1000; i++ {

		workers <- i

		go func(job int) {

			defer func() {
				<-workers
			}()

			time.Sleep(1 * time.Second)
			fmt.Println(job)
		}(i)

	}

	time.Sleep(2 * time.Second)

}
