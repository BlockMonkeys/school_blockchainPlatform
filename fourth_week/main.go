package main

import "fmt"

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 2 Token to NAMEONE")
	bc.AddBlock("Send 2 Token to NAMETWO")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
