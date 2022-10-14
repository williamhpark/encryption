package main

import (
	"fmt"

	"encryption-server/methods"
)

func main() {
	fmt.Println(methods.Caesar("test", 3, true))
	fmt.Println(methods.Caesar("qbpq", 3, false))
}
