package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	core "github.com/iden3/go-iden3-core"
)

func CreateClaim() {
	// set a long claim expiration data
	expirationDate := time.Date(2036, 10, 10, 0, 0, 0, 0, time.UTC)

	// A claim schema defines how a set of data must be stored inside a claim. In this example, we will use a
	// schema called KYCAgeCredential. According to this schema the birthday is stored in the first index slot
	// of the claim data structure, while the documentType is stored in the second data slot.
	// The hash of the schema is generated from the content of the schema document following the Claim Schema Generation Rules.
	// For our example, the hash of the schema is: 2e2d1c11ad3e500de68d7ce16a0a559e
	ageSchema, _ := core.NewSchemaHashFromHex("2e2d1c11ad3e500de68d7ce16a0a559e")

	// define the data slots
	bDay := big.NewInt(19960424)
	documentType := big.NewInt(1)

	// set revocation Nonce
	revocationNonce := uint64(1909830690)

	// set Id of the claim subject
	id, _ := core.IDFromString("113TCVw5KMeMp99Qdvub9Mssfz7krL9jWNvbdB7Fd2")

	// create the claim object
	claim, _ := core.NewClaim(
		ageSchema,
		core.WithExpirationDate(expirationDate),
		core.WithRevocationNonce(revocationNonce),
		core.WithIndexID(id),
		core.WithIndexDataInts(bDay, documentType),
	)

	claimData, _ := json.Marshal(claim)
	fmt.Println(string(claimData))
}
