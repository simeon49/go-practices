package demochain

import (
	"fmt"
	"log"
)

type Blockchain struct {
	Blocks []*Block
}

func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index-1 != oldBlock.Index {
		return false
	}
	if newBlock.PreBlockHash != oldBlock.Hash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func (bc *Blockchain) AppendBlock(newBlock *Block) {
	if !isValid(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
		log.Fatal("invalid block")
		return
	}
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("PreBlockHash: %s\n", block.PreBlockHash)
		fmt.Printf("Timestemp: %d\n", block.Timestamp)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Println()
	}
}

func (bc *Blockchain) SendData(data string) {
	newBlock := generateNewBlock(*bc.Blocks[len(bc.Blocks)-1], data)
	bc.AppendBlock(&newBlock)
}

func NewBlockchain() *Blockchain {
	bc := Blockchain{}
	genesisBlock := generateGenesisBlock()
	bc.Blocks = append(bc.Blocks, &genesisBlock)
	return &bc
}
