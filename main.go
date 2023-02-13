package main

import (
	"fmt"

	"github.com/iden3/go-iden3-crypto/babyjub"
)

func main() {
	privKey := babyjub.NewRandPrivKey()
	pubKey := privKey.Public()

	fmt.Println(pubKey)
}
