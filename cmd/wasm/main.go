package main

import (
	"fmt"
	"syscall/js"

	"github.com/innv8/webasm/cmd/wasm/wrappers"
)

func main() {
	fmt.Println("Go web assembly with fn")

	// load the json code
	js.Global().Set("formatJSON", wrappers.JsonWrapper())

	// make sure this is always running
	<-make(chan bool)
}
