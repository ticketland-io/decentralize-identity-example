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

	// Hereafter, this identity is represented as a mapping: ID => IdS. This gets published, together with all other
	// identities, inside the identities mapping, which is part of the State.sol contract. While the ID remains constant,
	// the Identity State will get updated as soon as the identity adds or revokes claims in its trees.
	//
	// No Personal Identifiable Information (PPI) is stored on-chain. From the IdS is impossible to retrieve any
	// information (represented as claim) stored inside the Identity Claims Tree
	id := identity.ID(state)

	fmt.Println("ID:", id)
}
