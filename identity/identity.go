package identity

import (
	"context"
	"fmt"

	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/go-merkletree-sql"
	"github.com/iden3/go-merkletree-sql/db/memory"
)

/// Each identity consists of Three Sparse Merkle Trees:
///
/// 1. ClT: A Claims tree that contains the claims issued by that particular identity
/// 2. ReT: A Revocation tree that contains the revocation nonces of the claims that have been revoked by that particular identity.
///         The revocation tree gets updated whenever an identity decides to revoke a claim. For instance, if a user decides to rotate
///         her keys, then she generates a key pair, creates a new authClaim with the public key from the key pair and adds the claim
///         to the Claims Tree. Now the user can revoke the old public key, so she adds an entry to the Revocation Tree with the claim
///         revocation nonce as an Index and zero as a Value.
/// 3. RoT: A Roots tree that contains the history of the tree roots from the Claims tree
///         The Roots Tree gets updated whenever the Identity Claims Tree root gets updated.
///
/// Claims issued by an identity are added to the Claims tree (we'll see in a while why that's not always the case).
/// The position of a claim inside the Sparse Merkle Tree is determined by the hash of the claim's Index while the
/// value stored inside the leaf will be the hash of the claim's Value.
///
/// An identity must issue at least one Auth Claim to operate properly. This is the first claim that is issued by an
/// identity and that must be added to the ClT.
func CreateIdentity(authClaim *core.Claim) {
	ctx := context.Background()
	
	// Create empty Claim's tree (Clt)
	clt, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32)

	// Create empty Revocation's tree (ReT)
	// The revocation tree gets updated whenever an identity decides to revoke a claim.
	ret, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32)

	// Create empty Root's tree (RoT)
	rot, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32)

	// Get the index and the value of the auth claim
	index, value, _ := authClaim.HiHv()

	// Add auth claim value to Claim's tree at the specific index
	clt.Add(ctx, index, value)

	// print the roots
	fmt.Println(clt.Root().BigInt(), ret.Root().BigInt(), rot.Root().BigInt())
}
