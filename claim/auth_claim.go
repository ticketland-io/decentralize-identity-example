package claim

import (
	"encoding/json"
	"fmt"

	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/go-iden3-crypto/babyjub"
)

// The most important building block of an identity is the Key Authorization Claim.
// This claim stores user's Baby Jubjub public key.
// An Auth Claim (https://docs.iden3.io/protocol/bjjkey/) must be included as a leaf inside the Identity Tree.
// All the actions performed by an Idenitity (such as claim issuance or revocation) require users to prove via
// a digital signature that they own the private key associated with the public key stored in the AuthClaim.
func AuthClaim(pubkey *babyjub.PublicKey) *core.Claim {
	authSchemHash, _ := core.NewSchemaHashFromHex("ca938857241db9451ea329256b9c06e5")

	// Add revocation nonce. Used to invalidate the claim. This may be a random number in the real implementation.
	revocationNonce := uint64(1)

	// Create the Auth Claim
	// According to the this schema (https://github.com/iden3/claim-schema-vocab/blob/main/schemas/json-ld/auth.json-ld),
	// the X and Y coordinate of the Baby Jubjub public key must be stored, respectively, in the first and second index data slot.
	authClaim, _ := core.NewClaim(
		authSchemHash,
		core.WithIndexDataInts(pubkey.X, pubkey.Y),
		core.WithRevocationNonce(revocationNonce),
	)

	authClaimData, _ := json.Marshal(authClaim)

	fmt.Println(string(authClaimData))

	return authClaim
}
