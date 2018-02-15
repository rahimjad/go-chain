package node

import (
	"../models"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var Address string = "q3nf394hjg-random-miner-address-34nf3i4nflkn3oi"
var PeerNodes []string
var BlockChain []*models.Block

// Private
func BlockHash(block *models.Block) []byte {
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
func CreateGenesisBlock() *models.Block {
	data := models.Data{ProofOfWork: 1, Transactions: []*models.Transaction{}}
	b := &models.Block{Index: 0, Timestamp: time.Now().Unix(), Data: data, PreviousHash: nil, Hash: nil}
	b.Hash = BlockHash(b)
	return b
}

func NextBlock(blockChain []*models.Block, lastBlock *models.Block) *models.Block {
	i := lastBlock.Index + 1
	data := models.Data{}

	b := &models.Block{Index: i, Timestamp: time.Now().Unix(), Data: data, PreviousHash: lastBlock.Hash, Hash: nil}
	b.Hash = BlockHash(b)

	return b
}

func GetBlocks() string {
	byteJSON, _ := json.Marshal(BlockChain)
	return string(byteJSON)
}

func FindNewChains() [][]*models.Block {
	var otherChains [][]*models.Block
	for _, url := range PeerNodes {
		resp, _ := http.Get(url + "/block")

		defer resp.Body.Close()

		jsonBlock, _ := ioutil.ReadAll(resp.Body)
		chain := make([]*models.Block, 0)
		json.Unmarshal(jsonBlock, &chain)

		otherChains = append(otherChains, chain)
	}
	return otherChains
}

func Consensus() {
	otherChains := FindNewChains()
	longestChain := BlockChain

	for _, chain := range otherChains {
		if len(longestChain) < len(chain) {
			longestChain = chain
		}
	}
	BlockChain = longestChain
}

func RegisterToBlockChain(b *models.Block) {
	BlockChain = append(BlockChain, b)
}

func ProofOfWork(lastProof int) int {
	v := 0
	for i := lastProof + 1; i%9 != 0 || i%lastProof != 0; i++ {
		v = i
	}
	return v
}
