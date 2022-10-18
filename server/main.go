package main

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"encryption-server/methods"
)

type Metadata struct {
	Message string `json:"message"`
	Key     string `json:"key"`
}

var privateKey rsa.PrivateKey

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

	var metadata Metadata
	// Decode the received body, store the metadata in `metadata`
	_ = json.NewDecoder(r.Body).Decode(&metadata)

	// Convert the Caesar Cipher key to an int
	key, err := strconv.Atoi(metadata.Key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Encrypt the message
	encryptedMessage, err := methods.CaesarEncrypt(metadata.Message, key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

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

	var metadata Metadata
	// Decode the received body, store the metadata in `metadata`
	_ = json.NewDecoder(r.Body).Decode(&metadata)

	// Convert the Caesar Cipher key to an int
	key, err := strconv.Atoi(metadata.Key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Decrypt the message
	decryptedMessage, err := methods.CaesarDecrypt(metadata.Message, key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

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

	var metadata Metadata
	// Decode the received body, store the metadata in `metadata`
	_ = json.NewDecoder(r.Body).Decode(&metadata)

	// Encrypt the message
	encryptedMessage, err := methods.AESEncrypt(metadata.Message, metadata.Key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

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

	var metadata Metadata
	// Decode the received body, store the metadata in `metadata`
	_ = json.NewDecoder(r.Body).Decode(&metadata)

	// Decrypt the message
	decryptedMessage, err := methods.AESDecrypt(metadata.Message, metadata.Key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Return the decrypted message to the client
	json.NewEncoder(w).Encode(decryptedMessage)
}

// Handler for encrypting a message using RSA
func rsaEncryptHandler(w http.ResponseWriter, r *http.Request) {
	// Error checking
	if r.URL.Path != "/rsa/encrypt" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	var metadata Metadata
	// Decode the received body, store the metadata in `metadata`
	_ = json.NewDecoder(r.Body).Decode(&metadata)

	// Generate a private key
	privateKey = methods.GenerateRSAPrivateKey()

	// Encrypt the message
	encryptedMessage, err := methods.RSAEncrypt(metadata.Message, privateKey.PublicKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Return the encrypted message to the client
	json.NewEncoder(w).Encode(encryptedMessage)
}

// Handler for decrypting a message using RSA
func rsaDecryptHandler(w http.ResponseWriter, r *http.Request) {
	// Error checking
	if r.URL.Path != "/rsa/decrypt" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	var metadata Metadata
	// Decode the received body, store the metadata in `metadata`
	_ = json.NewDecoder(r.Body).Decode(&metadata)

	// Decrypt the message
	decryptedMessage, err := methods.RSADecrypt(metadata.Message, privateKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Return the decrypted message to the client
	json.NewEncoder(w).Encode(decryptedMessage)
}

func main() {
	router := mux.NewRouter()

	/*
		Handler functions for request paths
	*/

	// Caesar Cipher
	router.HandleFunc("/caesar/encrypt", caesarEncryptHandler).Methods("POST")
	router.HandleFunc("/caesar/decrypt", caesarDecryptHandler).Methods("POST")

	// AES
	router.HandleFunc("/aes/encrypt", aesEncryptHandler).Methods("POST")
	router.HandleFunc("/aes/decrypt", aesDecryptHandler).Methods("POST")

	// RSA
	router.HandleFunc("/rsa/encrypt", rsaEncryptHandler).Methods("POST")
	router.HandleFunc("/rsa/decrypt", rsaDecryptHandler).Methods("POST")

	/*
		Serve the contents of the build directory that was produced as a part of `yarn run build` on the root `/`
	*/
	http.Handle("/", http.FileServer(http.Dir("./build")))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowedHeaders: []string{"*"},
	})
	handler := c.Handler(router)

	// Check if the port environment variable has been set and if so, use that, otherwise use a reasonable default
	port := os.Getenv("PORT")
	defaultPort := "8080"
	if port != "" {
		fmt.Printf("Starting server at port %v\n", port)
		log.Fatal(http.ListenAndServe(":"+port, handler))
	} else {
		fmt.Printf("Starting server at port %v\n", defaultPort)
		log.Fatal(http.ListenAndServe(":"+defaultPort, handler))
	}
}
