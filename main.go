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

	clt, ret, rot := identity.CreateIdentity(authClaim)
	// The very first identity state of an identity is defined as *Genesis State*
	state := identity.CreateIdentityState(clt, ret, rot)

	fmt.Println("identity state:", state)
}
