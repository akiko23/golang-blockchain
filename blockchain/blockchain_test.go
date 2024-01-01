package blockchain

import "testing"

func TestInitBlockChain(t *testing.T) {
	bc := InitBlockChain()

	wantLen := 1 // Only genesis block
	assertNumberOfBlocks(t, len(bc.Blocks), wantLen)
}

func TestAddBlock(t *testing.T) {
	bc := InitBlockChain()
	bc.AddBlock("Some test block")

	wantLen := 2 // Genesis block + the added one
	assertNumberOfBlocks(t, len(bc.Blocks), wantLen)
}

func assertNumberOfBlocks(t testing.TB, gotLen, wantLen int) {
	t.Helper()

	if gotLen != wantLen {
		t.Errorf("Expected %d blocks but got %d", wantLen, gotLen)
	}
}
