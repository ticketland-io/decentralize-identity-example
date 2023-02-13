package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/iden3/go-merkletree-sql"
	"github.com/iden3/go-merkletree-sql/db/memory"
)

func MerkleTree() {
	ctx := context.Background()

	// Tree Storage
	store := memory.NewMemoryStorage()

	// Merkle tree with 32 levels
	mt, _ := merkletree.NewMerkleTree(ctx, store, 32)

	// Add a lead to the tree with index 1 and value 10
	mt.Add(ctx, big.NewInt(1), big.NewInt(10))

	// Add another leaf
	mt.Add(ctx, big.NewInt(2), big.NewInt(15))

	// Proof of membership of a leaf with index 1
	proofExists, value, _ := mt.GenerateProof(ctx, big.NewInt(1), mt.Root())
	fmt.Println("Is included in the tree?:", proofExists.Existence)
	fmt.Println("The value is:", value)

	// Proof of non-membership of a leaf with index 4
	proofNotExists, _, _ := mt.GenerateProof(ctx, big.NewInt(4), mt.Root())
	fmt.Println("Proof of membership:", proofNotExists.Existence)
}
