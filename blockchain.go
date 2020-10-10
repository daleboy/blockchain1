package blockchain1

//Blockchain 区块链结构
type Blockchain struct {
	Blocks []*Block
}

//AddBlock 给区块链添加一个块
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewBlockchain 实现一个函数，来创建有创始区块的区块链
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
