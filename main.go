package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"

	"github.com/akiko23/golang-blockchain/blockchain"
)

func GetHash(b *blockchain.Block) []byte {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	return hash[:]
}

func main() {
	bc := blockchain.InitBlockChain()

	bc.AddBlock("First block after Genesis")
	bc.AddBlock("Second block after Genesis")
	bc.AddBlock("Third block after Genesis")

	for _, block := range bc.Blocks {
		fmt.Printf("Hash of the previous block: %x\n", block.PrevHash)
		fmt.Printf("Hash of the current block: %x\n", block.Hash)
		fmt.Printf("Data of the current block: %s\n\n", block.Data)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
