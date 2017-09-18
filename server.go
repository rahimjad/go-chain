package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"./node"
)

type Data struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

var nodeTransactions []*Data

// Handlers
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GO-CHAIN")
}

func handleTransaction(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var data Data

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

func listen() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/transaction", handleTransaction)
	http.ListenAndServe(":8080", nil)
}

func main() {
	node.RegisterToBlockChain(node.CreateGenesisBlock())
	listen()
}
