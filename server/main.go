package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"encryption-server/methods"
)

type Metadata struct {
	Message string `json:"message"`
	Key     string `json:"key"`
}

// Handler for encrypting a message using a Caesar Cipher
func caesarEncryptHandler(w http.ResponseWriter, r *http.Request) {
	// Error checking
	if r.URL.Path != "/caesar/encrypt" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var metadata Metadata
	// Decode the received body, store the metadata in `metadata`
	_ = json.NewDecoder(r.Body).Decode(&metadata)

	// Convert the Caesar Cipher key to an int
	key, err := strconv.Atoi(metadata.Key)
	if err != nil {
		log.Fatal(err)
	}

	// Encrypt the message
	encryptedMessage := methods.CaesarEncrypt(metadata.Message, key)

	// Return the encrypted message to the client
	json.NewEncoder(w).Encode(encryptedMessage)
}

// Handler for decrypting a message using a Caesar Cipher
func caesarDecryptHandler(w http.ResponseWriter, r *http.Request) {
	// Error checking
	if r.URL.Path != "/caesar/decrypt" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var metadata Metadata
	// Decode the received body, store the metadata in `metadata`
	_ = json.NewDecoder(r.Body).Decode(&metadata)

	// Convert the Caesar Cipher key to an int
	key, err := strconv.Atoi(metadata.Key)
	if err != nil {
		log.Fatal(err)
	}

	// Decrypt the message
	decryptedMessage := methods.CaesarDecrypt(metadata.Message, key)

	// Return the decrypted message to the client
	json.NewEncoder(w).Encode(decryptedMessage)
}

// Handler for encrypting a message using AES
func aesEncryptHandler(w http.ResponseWriter, r *http.Request) {
	// Error checking
	if r.URL.Path != "/aes/encrypt" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var metadata Metadata
	// Decode the received body, store the metadata in `metadata`
	_ = json.NewDecoder(r.Body).Decode(&metadata)

	// Encrypt the message
	encryptedMessage := methods.AESEncrypt(metadata.Message, metadata.Key)

	// Return the encrypted message to the client
	json.NewEncoder(w).Encode(encryptedMessage)
}

// Handler for decrypting a message using AES
func aesDecryptHandler(w http.ResponseWriter, r *http.Request) {
	// Error checking
	if r.URL.Path != "/aes/decrypt" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var metadata Metadata
	// Decode the received body, store the metadata in `metadata`
	_ = json.NewDecoder(r.Body).Decode(&metadata)

	// Decrypt the message
	decryptedMessage := methods.AESDecrypt(metadata.Message, metadata.Key)

	// Return the decrypted message to the client
	json.NewEncoder(w).Encode(decryptedMessage)
}

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

	// Handler functions for request paths

	// Caesar Cipher
	http.HandleFunc("/caesar/encrypt", caesarEncryptHandler)
	http.HandleFunc("/caesar/decrypt", caesarDecryptHandler)

	// AES
	http.HandleFunc("/aes/encrypt", aesEncryptHandler)
	http.HandleFunc("/aes/decrypt", aesDecryptHandler)

	fmt.Printf("Starting server at port 8080\n")
	// Tell the global HTTP server to listen for requests on port 8080
	// Pass `nil` for the `http.Handler` parameter, i.e. use the default server multiplexer
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
