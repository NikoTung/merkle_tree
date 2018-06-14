package main

import "fmt"

func main() {
	data := [][]byte{
		[]byte("node1"),
		[]byte("node2"),
		[]byte("node3"),
	}
	tree, _ := NewMerkleTree(data)
	c := make(chan []byte)

	go travel(tree, c)

	for v := range c {
		fmt.Printf("%x\n", v)
	}

}

func travel(tree *MerkleTree, c chan []byte) {

	walk(tree.Root, c)
	close(c)
}

func walk(node *MerkleNode, c chan []byte) {
	if node == nil {
		return
	}

	walk(node.Left, c)
	c <- node.Data
	walk(node.Right, c)
}
