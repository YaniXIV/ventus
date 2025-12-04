package main

import (
	"fmt"
	"ventus/merkle"
)

func main() {
	fmt.Println("Hello world!")
	
	// Example with SHA256 hasher
	hasher := &merkle.SHA256Hasher{}
	testTree := merkle.InitGSMT(hasher)
	
	// Example usage
	testData := []byte("test data")
	testTree.AddGSMT(testData)
	testTree.BuildGSMT()
	
	fmt.Println("Tree built successfully!")
	
	// Later you can use your Poseidon2 hasher like this:
	// poseidonHasher := &merkle.Poseidon2Hasher{}
	// poseidonTree := merkle.InitGSMT(poseidonHasher)
}
