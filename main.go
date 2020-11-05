package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	"rsc.io/quote"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) HashingBlock() {
	blockData := bytes.Join([][]byte{b.Hash, b.Data}, []byte{})
	blockHash := sha256.Sum256(blockData)
	b.Hash = blockHash[:]
}

func main() {
	fmt.Print(quote.Hello())
}
