//goroutines practice.
//goroutines are concurrent "threads" for go, which means functions can all run at the same time
//information found here: https://golangbot.com/goroutines/

package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello from goroutines")
}

//the goroutine returns immediately, so after executing, keyword "go" does not wait for the hello() function to execute

func main() {
	go hello()
	//stopping the execution of the main function allows goroutines to work.
	//there are many ways to stop the execution of the main function, and time.Sleep is one that illustrates the concept.
	//tutorial says using channels is the prod way of stopping main from running while goroutines finish up
	time.Sleep(1 * time.Second)
	fmt.Println("hello from main")
}
