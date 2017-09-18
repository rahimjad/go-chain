package node

import (
	"crypto/sha256"
	"time"
	"encoding/json"
	"log"
)

type Block struct {
	Index         int
	Timestamp     int64
	Data          map[string]interface{}
	PreviousHash  []byte
	Hash          []byte
}

var BlockChain []*Block
// Private
func blockHash(block *Block) []byte {
	h := sha256.New()

	dataBytes, err := json.Marshal(block.Data)
	if err != nil {
		log.Fatal(err)
		return []byte{}
	}

	blockString := string(block.Index) + string(block.Timestamp) + string(dataBytes) + string(block.PreviousHash)
	h.Write([]byte(blockString))

	return h.Sum(nil)
}

// Interface
func CreateGenesisBlock() *Block {
	b := &Block{Index: 0, Timestamp: time.Now().Unix(), Data: map[string]interface{}{}, PreviousHash: nil, Hash: nil}
	b.Hash = blockHash(b)
	return b
}

func NextBlock(blockChain []*Block, lastBlock *Block) *Block {
	i := lastBlock.Index + 1
	b := &Block{Index: i, Timestamp: time.Now().Unix(), Data: map[string]interface{}{}, PreviousHash: lastBlock.Hash, Hash: nil}
	b.Hash = blockHash(b)
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
