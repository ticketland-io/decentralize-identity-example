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
	expirationDate := time.Date(2361, 3, 22, 0, 44, 48, 0, time.UTC)

	// A claim schema defines how a set of data must be stored inside a claim. In this example, we will use a
	// schema called KYCAgeCredential (https://github.com/iden3/claim-schema-vocab/blob/main/schemas/json-ld/kyc-v2.json-ld). 
	// According to this schema the birthday is stored in the first index slot
	// of the claim data structure, while the documentType is stored in the second data slot.
	// The hash of the schema is generated from the content of the schema document following the Claim Schema Generation Rules
	// https://docs.iden3.io/getting-started/claim/claim-schema/. For our example, the hash of the schema is: 2e2d1c11ad3e500de68d7ce16a0a559e
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

	// The first 4 values of the claim represent the Index part of the claim while the last 4 represent the Value.
	// This is according to the claim structure defined here https://docs.iden3.io/protocol/claims-structure/
	//
	// If the subject is Self - identity i.e. The claim says something about itself, sections i_1, v_1 can be empty (0).
	// Claims are stored in the Merkle tree and the hash of the index slots ( hash(i_0,i_1,i_2,i_3) ) is unique for the whole tree.
	// This hash determines the leaf position where the value of the claim will be stored.
	// It means that you cannot have two claims with the same index inside the tree. As opposite to the index, the values slots could
	// be the same for different claims if their indexes are different.
	// Index:
	// {
	// 	"3613283249068442770038516118105710406958", // Claim Schema hash
	// 	"0", // ID Subject of the claim
	// 	"19960424", // First index data slot stores the date of birth
	// 	"1"  //  Second index data slot stores the document type
	// }
	//
	// Value:
	// { 
	// 	"227737944108667786680629310498", // Revocation nonce 
	// 	"0",
	// 	"0", // first value data slot
	// 	"0"  // second value data slot
	// }  
	fmt.Println(string(claimData))
}
