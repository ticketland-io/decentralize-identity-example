package identity

import "github.com/iden3/go-merkletree-sql"

/// The identity states are published on the blockchain under the identifier, anchoring the state of the identity
/// with the timestamp when it is published. In this way, the claims of the identity can be proved against the anchored
/// identity state at a certain timestamp. To transition from one state to the other, identities follow the transition functions.
///
/// The identity states can be published on the blockchain in one of the two ways: directly performing the transaction to
/// publish the root or indirectly using a Relay.
///
/// The Genesis State is the initial state of any identity, and does not need to be published on the blockchain, as the claims
/// under it can be verified against the identifier itself (that contains that identity state).
///
/// Every verification inside Iden3 protocol is executed against the Identity State. For instance, to prove the validity of
/// a specific claim issued by A to B (in case if the claims gets added to the claims tree):
///
/// 1. user B needs to produce a merkle proof of the existence of that claim inside user's A ClR
/// 2. user B needs to produce a merkle proof of non existence of the corresponding revocation nonce inside user's A ReT
func CreateIdentityState(clt, ret, rot *merkletree.MerkleTree) *merkletree.Hash {
	// An Identity State IdS is represented by the hash of the roots of these three merkle trees:
	// `IdS = H(ClR || ReR || RoR)`
	state, _ := merkletree.HashElems(
		clt.Root().BigInt(),
    ret.Root().BigInt(),
    rot.Root().BigInt(),
	)

	return state
}
