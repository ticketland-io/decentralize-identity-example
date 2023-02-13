package main

import (
	"did-example/claim"
	"fmt"

	"github.com/iden3/go-iden3-crypto/babyjub"
)

func main() {
	privKey := babyjub.NewRandPrivKey()
	pubKey := privKey.Public()

	fmt.Println(pubKey)

	// Merkle tree
	MerkleTree()
	claim.CreateClaim()
	claim.AuthClaim(pubKey)
}
