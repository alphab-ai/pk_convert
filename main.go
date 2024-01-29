package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/crypto"
)

// REPLACE THIS STRING WITH YOUR PRIVATE KEY 
var PRIVKEY = "REPLACE THIS STR WITH YOUR PRIVATE KEY"

// passphrase to decrypt your key
// if you plan to import privkey to a binary 8+ lenght pwd required
var PASSWORD = "12345678"

func main() {

	privKeyBytes, err := hex.DecodeString(PRIVKEY)
	
	if err != nil {
		fmt.Println("Error decoding hex private key:", err)
		os.Exit(1)
	}

	privKey := &secp256k1.PrivKey{
		Key: privKeyBytes,
	}

	armoredString := crypto.EncryptArmorPrivKey(privKey, PASSWORD, "secp256k1")

	fmt.Println("Armored Private Key:\n")
	fmt.Println(armoredString)
}