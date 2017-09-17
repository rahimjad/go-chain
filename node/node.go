package node

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	index         int
	timestamp     int64
	data          string
	previous_hash []byte
	hash          []byte
}

var BlockChain []*Block
// Private
func blockHash(block *Block) []byte {
	h := sha256.New()
	blockString := string(block.index) + string(block.timestamp) + block.data + string(block.previous_hash)
	h.Write([]byte(blockString))
	return h.Sum(nil)
}

// Interface
func CreateGenesisBlock() *Block {
	b := &Block{index: 0, timestamp: time.Now().Unix(), data: "Genesis Block", previous_hash: nil, hash: nil}
	b.hash = blockHash(b)
	return b
}

func NextBlock(blockChain []*Block, lastBlock *Block) *Block {
	i := lastBlock.index + 1
	b := &Block{index: i, timestamp: time.Now().Unix(), data: fmt.Sprintf("Hey I am block %d", i), previous_hash: lastBlock.hash, hash: nil}
	b.hash = blockHash(b)
	return b
}

func RegisterToBlockChain(b *Block) {
	BlockChain = append(BlockChain, b)
}

func ProofOfWork(lastProof int) int {
  v := 0
  for i := lastProof + 1; i % 9 != 0 || i % lastProof != 0; i++ {
    v = i
  }
  return v
}
