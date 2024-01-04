package blockchain

import (
	badger "github.com/dgraph-io/badger"
)

type BlockChainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

func (bci *BlockChainIterator) Next() *Block {
	var blockBytes []byte

	err := bci.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(bci.CurrentHash)
		Handle(err)

		err = item.Value(func(val []byte) error {
			blockBytes = append(blockBytes, val...)
			return nil
		})

		return err
	})
	Handle(err)

	block := Deserialize(blockBytes)
	bci.CurrentHash = block.PrevHash

	return block
}
