package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Data struct {
	from   string
	to     string
	amount float64
}
type Block struct {
	index         int
	timestamp     int64
	data          string
	previous_hash []byte
	hash          []byte
}

var blockChain []*Block

func blockHash(block *Block) []byte {
	h := sha256.New()
	blockString := string(block.index) + string(block.timestamp) + block.data + string(block.previous_hash)
	h.Write([]byte(blockString))
	return h.Sum(nil)
}

func createGenesisBlock() *Block {
	b := &Block{index: 0, timestamp: time.Now().Unix(), data: "Genesis Block", previous_hash: nil, hash: nil}
	b.hash = blockHash(b)
	registerToBlockChain(b)
	return b
}

func nextBlock(lastBlock *Block) *Block {
	i := lastBlock.index + 1
	b := &Block{index: i, timestamp: time.Now().Unix(), data: fmt.Sprintf("Hey I am block %d", i), previous_hash: lastBlock.hash, hash: nil}
	b.hash = blockHash(b)
	registerToBlockChain(b)
	return b
}

func registerToBlockChain(b *Block) {
	blockChain = append(blockChain, b)
}

func proofOfWork(lastProof int) int {
  v := 0
  for i := lastProof + 1; i % 9 != 0 || i % lastProof != 0; i++ {
    v = i
  }
  return v
}

func main() {
	fmt.Println(blockChain)
	b := createGenesisBlock()
	fmt.Printf("%+v\n", b)
	fmt.Println(blockChain)
	fmt.Println("%+v\n", nextBlock(b))
	fmt.Println(blockChain)
}
