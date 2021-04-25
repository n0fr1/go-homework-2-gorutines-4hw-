package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	sigs := make(chan os.Signal)
	done := make(chan bool)

	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		fmt.Printf("%s", <-sigs)
		close(done) //close channel or done <- true
	}()

	fmt.Printf("%s \n", "awaiting signal")

	<-done //trying to read
	time.Sleep(1 * time.Second)
	fmt.Printf("\n%s", "exiting")
}
