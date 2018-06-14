package main

import (
	"crypto/sha256"
	"fmt"
)

type MerkleTree struct {
	Root *MerkleNode
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

type ErrMerkleNodeData []byte

func (e ErrMerkleNodeData) Error() string {
	return fmt.Sprintf("node can not be empty")
}

func NewMerkleTree(data [][]byte) (*MerkleTree, error) {

	if len(data) <= 0 {
		return nil, ErrMerkleNodeData{}
	}

	var nodes []MerkleNode

	for _, b := range data {
		node := NewMerkleNode(nil, nil, b)
		nodes = append(nodes, *node)
	}

	for len(nodes) > 1 {
		var n []MerkleNode
		for i := 0; i < len(nodes); i += 2 {
			if i+1 <= len(nodes)-1 {
				node := NewMerkleNode(&nodes[i], &nodes[i+1], nil)
				n = append(n, *node)
			} else {
				node := NewMerkleNode(&nodes[i], &nodes[i], nil)
				n = append(n, *node)
			}
		}

		nodes = n[:]
	}

	return &MerkleTree{&nodes[0]}, nil
}

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	merkleNode := MerkleNode{}

	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		merkleNode.Data = hash[:]
	} else {
		preHash := append(left.Data, right.Data...)
		hash := sha256.Sum256(preHash)
		merkleNode.Data = hash[:]
	}
	merkleNode.Left = left
	merkleNode.Right = right

	return &merkleNode
}
