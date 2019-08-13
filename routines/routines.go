//goroutines practice.
//goroutines are concurrent "threads" for go, which means functions can all run at the same time
//information found here: https://golangbot.com/goroutines/

package main

import (
	"fmt"
)

func hello() {
	fmt.Println("hello from goroutines")
}

//trying to execute the goroutine immediately like this will not work.
//the goroutine returns immediately, so after executing, keyword "go" does not wait for the hello() function to execute

func main() {
	go hello()
	fmt.Println("hello from main")
}
