package main

import (
	"fmt"
	"go.bergrend.dev/mod"
)

func main() {
	fmt.Println("[VERSION] v0.1.1")
	fmt.Println("cmd mod start")
	mod.Go()
	fmt.Println("cmd mod end")
}
