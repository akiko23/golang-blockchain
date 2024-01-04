package blockchain

import (
	badger "github.com/dgraph-io/badger"
)

const (
	dbPath = "./tmp/blocks"
)

type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

func InitBlockChain() *BlockChain {
	opt := badger.DefaultOptions(dbPath)

	db, err := badger.Open(opt)
	Handle(err)

	lastHash, err := getLastHash(db)
	Handle(err)

	bc := &BlockChain{lastHash, db}
	return bc
}

func (bc *BlockChain) AddBlock(data string) {
	lastHash, err := getLastHash(bc.Database)
	Handle(err)

	newBlock := CreateBlock(data, lastHash)

	err = addBlockToDb(bc, newBlock)
	Handle(err)

	// Update block chain last hash
	bc.LastHash = newBlock.Hash
}

func (bc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{bc.LastHash, bc.Database}
}
