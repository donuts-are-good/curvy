package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"log"

	"golang.org/x/crypto/curve25519"
)

// GenerateEd25519Keys takes no params and returns
// private key, and public key, and error
func GenerateEd25519Keys() ([]byte, []byte, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		log.Println(err)
	}
	return privateKey, publicKey, err
}

// GenerateX25519Keys takes an ed25519 private key
// and generates two [32]byte for private and public
// x25519 keys using curve25519.X25519
func GenerateX25519Keys(ed25519PrivateKey ed25519.PrivateKey) ([32]byte, [32]byte, error) {
	var privateKey [32]byte
	var publicKey [32]byte
	copy(privateKey[:], ed25519PrivateKey[:32])
	curve25519.ScalarBaseMult(&publicKey, &privateKey)

	x25519PrivateKey, err := curve25519.X25519(privateKey[:], publicKey[:])
	if err != nil {
		return [32]byte{}, [32]byte{}, err
	}
	copy(privateKey[:], x25519PrivateKey)
	return privateKey, publicKey, nil
}

func main() {

	// Generate a set of ed25519 keys
	ed25519PrivateKey, ed25519PublicKey, errEd25519 := GenerateEd25519Keys()
	if errEd25519 != nil {
		log.Println("Error generating ed25519 keys: ", errEd25519)
		return
	}

	// Derive a set of x25519 keys
	x25519PrivateKey, x25519PublicKey, errX25519 := GenerateX25519Keys(ed25519PrivateKey)
	if errX25519 != nil {
		log.Println("Error generating x25519 keys: ", errX25519)
	}

	// Encode these keys to strings
	ed25519PrivateKeyHex := hex.EncodeToString(ed25519PrivateKey)
	ed25519PublicKeyHex := hex.EncodeToString(ed25519PublicKey)
	x25519PrivateKeyHex := hex.EncodeToString(x25519PrivateKey[:])
	x25519PublicKeyHex := hex.EncodeToString(x25519PublicKey[:])

	// Trim the stupid one that has the appended bit.
	// Show the goods as strings.
	fmt.Println("(ed25519) Private Key\t", ed25519PrivateKeyHex[:64])
	fmt.Println("(ed25519) Public Key\t", ed25519PublicKeyHex)
	fmt.Println("(x25519)  Private Key\t", x25519PrivateKeyHex)
	fmt.Println("(x25519)  Public Key\t", x25519PublicKeyHex)

}
