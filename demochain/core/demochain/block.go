package demochain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index        int64  // 区块编号
	Timestamp    int64  // 时间戳
	PreBlockHash string // 前一个区块链的hash
	Hash         string // 当前hash

	Data string // 区块数据
}

func calculateHash(b Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + b.PreBlockHash + b.Data
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

func generateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.PreBlockHash = preBlock.Hash
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

func generateGenesisBlock() Block {
	genesisBlock := Block{}
	genesisBlock.Index = 0
	genesisBlock.Timestamp = time.Now().Unix()
	genesisBlock.PreBlockHash = ""
	genesisBlock.Data = ""
	genesisBlock.Hash = calculateHash(genesisBlock)
	return genesisBlock
}
