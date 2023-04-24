package BLC

type BlockChain struct {
	Blocks []*Block // 存储所有区块
}

// 增加区块
func (blc *BlockChain) AddBlockChain(data string, height int64, preHash []byte) {
	newblock := NewBlock(data, height, preHash)
	blc.Blocks = append(blc.Blocks, newblock)
}

// 1.创建带有创世区块的区块链
func CreateBlockWithGennesisBlock() *BlockChain {
	genesisBlock := CreateGenesisBlock("Genesis Data..")

	return &BlockChain{[]*Block{genesisBlock}}

}
