package merkle

import (
	"crypto/sha256"
	"math"
)

// tree -> layers -> leafs
type leaf = []byte
type Layer = []leaf
type Tree = []Layer

type GSMerkleTree struct {
	Root   leaf
	Levels Tree
	Hasher Hasher
}

type Hasher interface {
	Hash(data []byte) []byte
}

type SHA256Hasher struct{}
type Poseidon2Hasher struct{
	BinaryPath string
}
type Option func(*Poseidon2Hasher)


func WithBinaryPath(path string)Option{
	return func(p *Poseidon2Hasher){
		p.BinaryPath = path
	}
}

func NewPoseidon2Hasher(opts ...Option) *Poseidon2Hasher {
    h := &Poseidon2Hasher{
        BinaryPath: "/usr/bin/poseidon2", // default
    }

    // Apply any options
    for _, opt := range opts {
        opt(h)
    }

    return h
}

func (*SHA256Hasher) Hash(data []byte) []byte {
	hasher := sha256.New()
	hasher.Write(data)
	hashSum := hasher.Sum(nil)

	return hashSum[:]
}

func (*Poseidon2Hasher) Hash(data []byte) []byte{
	

}

// Example: To add a Poseidon2Hasher, implement the Hasher interface:
//
//
// func (p *Poseidon2Hasher) Hash(data []byte) []byte {
//     // Call your Rust PoseidonHash function via CGO
//     // Convert result to []byte
//     // return hashResult
// }
//
// Then use it: tree := InitGSMT(&Poseidon2Hasher{})

// Tree[ Layer[ Hash[] ] ]
// Initialize a new empty tree with the provided hasher.
// This function is hash-agnostic and accepts any implementation of the Hasher interface.
func InitGSMT(hasher Hasher) *GSMerkleTree {
	// initialize an empty merkle tree with one empty layer
	return &GSMerkleTree{
		Root:   nil,
		Levels: Tree{Layer{}},
		Hasher: hasher,
	}
}

// Init_GSMT_SHA256 is a convenience function that initializes a tree with SHA256 hasher
// Deprecated: Use InitGSMT(&SHA256Hasher{}) instead for consistency
func Init_GSMT_SHA256(h Hasher) *GSMerkleTree {
	return InitGSMT(h)
}

// InitGSMTWithSHA256 is a convenience function that initializes a tree with SHA256 hasher
func InitGSMTWithSHA256() *GSMerkleTree {
	return InitGSMT(&SHA256Hasher{})
}

func (t *GSMerkleTree) AddGSMT(l leaf) {
	result := t.Hasher.Hash(l)
	t.Levels[0] = append(t.Levels[0], result)
}

func (t *GSMerkleTree) BuildGSMT() {
	if len(t.Levels) == 0 || len(t.Levels[0]) == 0 {
		return
	}

	currentLayer := t.Levels[0]
	
	// Build layers until we have a root
	for len(currentLayer) > 1 {
		var nextLayer Layer
		
		// Pair leaves: if even index, hash with next; if odd, hash with previous
		// Handle odd number of leaves by duplicating the last one
		for i := 0; i < len(currentLayer); i += 2 {
			var pairHash []byte
			
			if i+1 < len(currentLayer) {
				// Pair exists: hash left and right together
				combined := append(currentLayer[i], currentLayer[i+1]...)
				pairHash = t.Hasher.Hash(combined)
			} else {
				// Odd number: duplicate the last leaf and hash with itself
				combined := append(currentLayer[i], currentLayer[i]...)
				pairHash = t.Hasher.Hash(combined)
			}
			
			nextLayer = append(nextLayer, pairHash)
		}
		
		// Add the new layer to the tree
		t.Levels = append(t.Levels, nextLayer)
		currentLayer = nextLayer
	}
	
	// Set the root
	if len(t.Levels) > 0 && len(t.Levels[len(t.Levels)-1]) > 0 {
		t.Root = t.Levels[len(t.Levels)-1][0]
	}
}

func (t *GSMerkleTree) CalcDepth() int {
	k := math.Ceil(math.Log2(float64(len(t.Levels[0]))))
	return int(k)
}
