package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	// 高度
	Height int64
	// 上一个区块哈希
	PreBlockHash []byte
	// 哈希
	Hash []byte
	// 交易数据
	Data []byte
	// 时间戳
	Timestamp int64
}

func (block *Block) SetHash() {
	// height --->[]byte
	heightByte := IntToHex(block.Height)
	// 时间戳转换成数组
	timeStr := strconv.FormatInt(block.Timestamp, 2)
	timeByte := []byte(timeStr)

	// 拼接属性
	blockByte := bytes.Join([][]byte{heightByte, block.PreBlockHash, block.Data, timeByte}, []byte{})

	// 生成hash
	hash := sha256.Sum256(blockByte) // 32位数组
	block.Hash = hash[:]
}
func NewBlock(data string, height int64, preBlockHash []byte) *Block {
	block := &Block{height, preBlockHash, nil, []byte(data), time.Now().Unix()}
	block.SetHash()
	return block
}

// 创世区块
func CreateGenesisBlock(data string) *Block {
	pre := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	block := &Block{1, pre, nil, []byte(data), time.Now().Unix()}
	block.SetHash()
	return block
}
