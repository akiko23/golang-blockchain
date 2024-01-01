package blockchain

type BlockChain struct {
	Blocks []*Block
}

func InitBlockChain() *BlockChain {
	bc := &BlockChain{[]*Block{GenesisBlock()}}
	return bc
}

func (bc *BlockChain) AddBlock(data string) {
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := CreateBlock(data, lastBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}
