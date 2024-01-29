package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"flag"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/crypto"
)

// REPLACE THIS STRING WITH YOUR PRIVATE KEY 
// var PRIVKEY = "REPLACE THIS STR WITH YOUR PRIVATE KEY"

// passphrase to decrypt your key
// if you plan to import privkey to a binary 8+ lenght pwd required
// var PASSWORD = "12345678"

func main() {
	pk := flag.String("pk", " ", "hex private key 64 chars")
    password := flag.String("password", "12345678", "")
    flag.Parse()

	privKeyBytes, err := hex.DecodeString(*pk)
	
	if err != nil {
		fmt.Println("Error decoding hex private key:", err)
		os.Exit(1)
	}

	privKey := &secp256k1.PrivKey{
		Key: privKeyBytes,
	}

	armoredString := crypto.EncryptArmorPrivKey(privKey, *password, "secp256k1")

	// fmt.Println("Armored Private Key:\n")
	fmt.Println(armoredString)
}