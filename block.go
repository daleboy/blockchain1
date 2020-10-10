package blockchain1

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

//Block 包含关键信息的区块结构,简化版
type Block struct {
	Timestamp     int64  //当前时间戳
	Data          []byte //区块存储的实际有效信息
	PrevBlockHash []byte //存储前一个块的哈希
	Hash          []byte //当前块的哈希
}

//SetHash 计算区块的哈希并设置到区块的哈希变量
//连接一个区块的哈希与区块存储的实际有效信息以及区块的时间戳的二进制码后，
//对其进行哈希计算得到当前区块的哈希
//注意，这里对Block类型定义的一个方法
func (b *Block) SetHash() {
	//整形不可以直接转为二进制，需要先转为字符串
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10)) //先转为10进制然后以字符串形式返回后强制转为二进制
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

//NewBlock 创建一个普通区块,返回指针
func NewBlock(data string, prevBlockHash []byte) *Block {
	//time.Now()返回Time类型，Unix()返回int64值
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}

	block.SetHash()
	return block
}

// NewGenesisBlock 创建创始区块
func NewGenesisBlock() *Block {
	return NewBlock("Genisis Block", []byte{})
}
