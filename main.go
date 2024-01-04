package main

import (
	"github.com/akiko23/golang-blockchain/blockchain"
)

func main() {
	bc := blockchain.InitBlockChain()
	defer bc.Database.Close()

	cli := CommandLine{bc}
	cli.run()
}
