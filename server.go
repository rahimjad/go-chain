package main

import (
	"./models"
	"./node"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var transactions []*models.Transaction

// Handlers
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GO-CHAIN")
}

func handleTransaction(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var data models.Transaction

	err := dec.Decode(&data)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("From: %s\n", data.From)
	fmt.Printf("To: %s\n", data.To)
	fmt.Printf("Amount: %f\n", data.Amount)

	output, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func handleMine(w http.ResponseWriter, r *http.Request) {
	node.Consensus()
	// Get last Proof
	lastBlock := node.BlockChain[len(node.BlockChain)-1]
	lastProof := lastBlock.Data.ProofOfWork

	// Compute new proof
	proof := node.ProofOfWork(lastProof)

	// Create transaction as reward for computing new proof
	newTransaction := &models.Transaction{From: "network", To: node.Address, Amount: 1}

	// Append to current transactions
	transactions = append(transactions, newTransaction)

	// build new Data
	data := models.Data{ProofOfWork: proof, Transactions: transactions}

	// Create new block
	newBlock := &models.Block{}
	newBlock.Data = data
	newBlock.Index = lastBlock.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.PreviousHash = lastBlock.Hash
	newBlock.Hash = node.BlockHash(lastBlock)

	// Append to blockchain
	node.BlockChain = append(node.BlockChain, newBlock)

	// Clear out transactions
	transactions = transactions[:0]

	// output
	output, err := json.Marshal(newBlock)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func handleGetBlocks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(node.GetBlocks()))
}

func listen() {
	port := ":8080"
	http.HandleFunc("/", handler)
	http.HandleFunc("/transaction", handleTransaction)
	http.HandleFunc("/mine", handleMine)
	http.HandleFunc("/block", handleGetBlocks)
	fmt.Println("Content served on", port)
	http.ListenAndServe(port, nil)
}

func main() {
	node.PeerNodes = append(node.PeerNodes, "http://localhost:8080")
	node.RegisterToBlockChain(node.CreateGenesisBlock())
	listen()
}
