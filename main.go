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

func main() {

	privKeyBytes, err := hex.DecodeString(PRIVKEY)
	
	if err != nil {
		fmt.Println("Error decoding hex private key:", err)
		os.Exit(1)
	}

	privKey := &secp256k1.PrivKey{
		Key: privKeyBytes,
	}

	armoredString := crypto.EncryptArmorPrivKey(privKey, "12345678", "secp256k1")

	fmt.Println("Armored Private Key:\n")
	fmt.Println(armoredString)
}