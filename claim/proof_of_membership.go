package claim

import (
	"fmt"
	"os"

	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/go-iden3-crypto/keccak256"
)

// The ProofOfDaoMembership claim should attest that a person covers a role inside a specific DAO.
// Information such as the identifier of the DAO or the identifier of the subject of the claim don't
// need to be encoded inside one of the four data slots allocated for claim information (i_2,i_3, v_2, v_3):
//
// 	1. The information about the specific DAO can be inferred from the claim issuer identifier
// 	2. The information about the individual subject of the claim is already stored in the i_1 or v_1 data slot of the claim
//
// A further information that must be included in the claim is the role of the individual inside a DAO.
// This will be the added inside one of the data slots (i_2,i_3,v_2,v_3).
//
// Decide where to store this information, should it be inside index data slots or value data slots?
// Claim's index determines its uniqueness inside the issuer's claims tree. There cannot be more than one claim with the same index.
// In this case, the question is whether to store the information with type role inside i_2 or v_2.
//
// 1. Storing the role inside i_2 means that the uniqueness inside the tree is determined by the combination "person identifier + role"
// 2. Storing the role inside v_2 means that the uniqueness inside the tree is only determined the person identifier
//
// DAO member covers more than one role, it makes more sense to store the role inside i_2.
func ProofOfMembership() {
	schema, _ := os.ReadFile("../schema/proof-of-dao-membership.json-ld")
	
	var schemHash core.SchemaHash
	h := keccak256.Hash(schema, []byte("ProofOfDaoMembership"))
	
	// take the slast 16 bytes https://docs.iden3.io/protocol/claim-schema/#schema-hash
	copy(schemHash[:], h[len(h) - 16:])
	
	hashHex, _ := schemHash.MarshalText()
	fmt.Println(string(hashHex))
}
