package blockchain

import (
	"fmt"

	badger "github.com/dgraph-io/badger"
)

func getLastHash(bcDatabase *badger.DB) ([]byte, error) {
	var (
		lastHash    []byte
		lastHashKey []byte
	)
	lastHashKey = []byte("lh")

	err := bcDatabase.Update(func(txn *badger.Txn) error {
		_, err := txn.Get(lastHashKey)
		if err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found")
			genesis := GenesisBlock()

			err = txn.Set(genesis.Hash, genesis.Serialize())
			Handle(err)

			err = txn.Set(lastHashKey, genesis.Hash)
			Handle(err)

			lastHash = append([]byte{}, genesis.Hash...)
		} else {
			item, err := txn.Get(lastHashKey)
			Handle(err)

			err = item.Value(func(val []byte) error {
				lastHash = append([]byte{}, val...)
				return err
			})
		}
		return err
	})
	return lastHash, err
}

func addBlockToDb(bc *BlockChain, newBlock *Block) error {
	err := bc.Database.Update(func(txn *badger.Txn) error {
		// Add new kv where the key is block's hash and value is serialized block
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		Handle(err)

		// Update last hash
		err = txn.Set([]byte("lh"), newBlock.Hash)
		return err
	})
	return err
}
