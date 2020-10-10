package blockchain1

//Blockchain 区块链结构
type Blockchain struct {
	blocks []*Block
}

//AddBlock 给区块链添加一个块
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain 实现一个函数，来创建有创始区块的区块链
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
