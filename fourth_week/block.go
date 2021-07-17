package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block 구조체
// Block Header 구조형성
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

//Block 구조체를 사용해 객체를 만들어서 블록을 만듬
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{Timestamp: time.Now().Unix(), Data: []byte(data), PrevBlockHash: prevBlockHash, Hash: []byte{}}
	block.SetHash()

	return block
}

// Block Header를 HASH 알고리즘으로 돌려서 HASH 도출.
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}
