package main

import (
	"did-example/claim"
	"did-example/identity"
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
	authClaim := claim.AuthClaim(pubKey)
	claim.ProofOfMembership()

	identity.CreateIdentity(authClaim)
}
