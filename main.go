package main

import (
	"fmt"
	"strconv"
)

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send more BTC to Ivan")

	for _, block := range bc.blocks {
		pow := NewProofOfWork(block)
		fmt.Printf("Row: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

}
