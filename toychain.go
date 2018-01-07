package main

import (
	"crypto/sha256"
	"math/rand"
	"time"
	"bytes"
	"strconv"
	"fmt"
)

const BLOCKCHAIN_SIZE = 4


type BlockChain []Block

func (bc BlockChain) VerifyChain() {
	println(fmt.Sprintf("Verifying chain..."))
	var previousHash []byte
	var nerrors = 0
	for i, b := range bc {
		if i > 0 {
			if bytes.Compare(previousHash, b.PreviousHashValue) != 0 {
				println(fmt.Sprintf("Verification failed on node %v", i))
				nerrors += 1
			} else {
				println(fmt.Sprintf("Verification passed on node %v", i))
			}
		}
		previousHash = b.ComputeHash()
	}
	println(fmt.Sprintf("Verification finished with %v errors", nerrors))
}


func (bc BlockChain) AddBlock(b Block) BlockChain {
	if len(bc) > 0 {
		pb := bc[len(bc)-1]
		b.PreviousHashValue = pb.HashValue
		b.HashValue = b.ComputeHash()
		println(fmt.Sprintf("Added block %v", b))
	}
	return append(bc, b)
}


func DummyBlock() Block {
	data := make([]byte, 6)
	rand.Read(data)
	phash := make([]byte, 32)
	rand.Read(phash)
	ts := time.Now().Nanosecond()
	b := Block{Timestamp:ts, Data: data, PreviousHashValue: phash}
	hash := b.ComputeHash()
	b.HashValue = hash
	return b
}


type Block struct {
	Timestamp int
	Data []byte
	HashValue []byte
	PreviousHashValue []byte
}


func (b *Block) ComputeHash() []byte {
	h256 := sha256.New()

	tsBytes := []byte(strconv.Itoa(b.Timestamp))
	bb := make([]byte, 0)
	bb = append(bb, tsBytes...)
	bb = append(bb, b.Data...)
	bb = append(bb, b.PreviousHashValue...)

	h256.Write(bb)
	return h256.Sum(nil)
}


func main() {

	bc := make(BlockChain, 0)
	for i := 0; i < BLOCKCHAIN_SIZE; i++ {
		bc = bc.AddBlock(DummyBlock())
		time.Sleep(time.Second)
	}
	bc.VerifyChain()
}