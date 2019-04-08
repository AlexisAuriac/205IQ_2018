package main

import (
	"fmt"
	"os"
)

func main() {
	var params *Params = ParseArgv(os.Args[1:])

	fmt.Println("u:", params.u)
	fmt.Println("s:", params.s)
	fmt.Println("iq1:", params.iq1)
	fmt.Println("iq2:", params.iq2)
}
