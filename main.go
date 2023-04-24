package main

import (
	"blockchain/BLC"
	"fmt"
)

func main() {
	// block := BLC.NewBlock("test block", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	// fmt.Println(block)

	genesisBlockChain := BLC.CreateBlockWithGennesisBlock()
	fmt.Println(genesisBlockChain.Blocks[0])

	genesisBlockChain.AddBlockChain("test", 2, genesisBlockChain.Blocks[0].Hash)
	fmt.Println(genesisBlockChain.Blocks[1])
}
