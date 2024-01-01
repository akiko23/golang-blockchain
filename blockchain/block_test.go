package blockchain

import (
	"bytes"
	"crypto/sha256"
	"reflect"
	"testing"
)

func getBlockHash(block *Block) []byte {
	info := bytes.Join([][]byte{block.Data, block.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	return hash[:]
}

func TestCreateBlock(t *testing.T) {
	data := "Test Block"
	prevHash := sha256.Sum256([]byte(data))

	block := CreateBlock(data, prevHash[:])
	block.DeriveHash()

	got := block.Hash
	want := getBlockHash(block)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Want %x but got %x", want, got)
	}
}

func TestGenesisBlock(t *testing.T) {
	gb := GenesisBlock()

	got := gb.PrevHash
	want := []byte{}

	if !reflect.DeepEqual(gb.PrevHash, want) {
		t.Errorf("Wanted empty PrevHash for genesis block (%x), but got %x", want, got)
	}
}
