package main

import (
	"fmt"

	"encryption-server/methods"
)

func main() {
	fmt.Println(methods.Caesar("test", 3, true))
	fmt.Println(methods.Caesar("qbpq", 3, false))

	fmt.Println(methods.AES("This is a secret", "thisis32bitlongpassphraseimusing", true))
	fmt.Println(methods.AES("145149d64a1a3c4025e67665001a3167", "thisis32bitlongpassphraseimusing", false))
}
