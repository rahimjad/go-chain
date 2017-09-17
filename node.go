package main

import (
	"crypto/sha256"
	"fmt"
	"time"
	"./server"
	"encoding/base64"
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
	blockChain = append(blockChain, createGenesisBlock())
	previousBlock := blockChain[0]
	numOfBlockToAdd := 20

	for i := 0; i < numOfBlockToAdd; i++ {
		blockToAdd := nextBlock(previousBlock)
		blockChain = append(blockChain, blockToAdd)
		fmt.Printf("Block %d has been added to the blockchain!\n", blockToAdd.index)
		fmt.Printf("Hash: %s\n", base64.URLEncoding.EncodeToString(blockToAdd.hash))
		previousBlock = blockToAdd
	}

	server.Listen()
}
