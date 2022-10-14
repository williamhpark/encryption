package main

import (
	"fmt"

	"encryption-server/methods"
)

func main() {
	fmt.Println("CAESAR CIPHER")
	secretMessage := "This is super secret message!"
	encryptedMessage := methods.CaesarEncrypt(secretMessage, 3)
	fmt.Println(encryptedMessage)
	fmt.Println(methods.CaesarDecrypt(encryptedMessage, 3))
	fmt.Println()

	fmt.Println("AES")
	secretMessage = "This is a secret"
	encryptedMessage = methods.AESEncrypt(secretMessage, "thisis32bitlongpassphraseimusing")
	fmt.Println(encryptedMessage)
	fmt.Println(methods.AESDecrypt(encryptedMessage, "thisis32bitlongpassphraseimusing"))
	fmt.Println()

	fmt.Println("RSA")
	privateKey := methods.GenerateRSAPrivateKey()
	publicKey := methods.GenerateRSAPublicKey(&privateKey)
	secretMessage = "This is super secret message!"
	encryptedMessage = methods.RSAEncrypt(secretMessage, publicKey)
	fmt.Println(encryptedMessage)
	fmt.Println(methods.RSADecrypt(encryptedMessage, privateKey))
}
