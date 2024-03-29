package main

import (
	"fmt"
	"github.com/jase-d-ace/httpreqs"
	"github.com/jase-d-ace/stringutil"
)

//this is the main brain of a golang program.
//since the package is called "main," this file will be generate an executable that you use in your terminal
//the executable then runs the main function of this file. In this case, it runs two simple functions that have been imported from other folders in this workspace.

type Example struct {
	Foo, Bar, Boo string
}

var exArray = []Example{
	Example{"something", "else", "yeah"},
	Example{"no", "yes", "maybe"},
}

func printArray(arr []Example) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	httpreqs.MakeReq()
	printArray(exArray)
}
