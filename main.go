package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

var Blockchain []Block

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	hash := sha256.New()
	hash.Write([]byte(record))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.Hash = calculateHash(newBlock)
	newBlock.PrevHash = oldBlock.Hash

	return newBlock, nil
}

func isBlockValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index != oldBlock.Index+1 {
		return false
	}

	if newBlock.Hash != calculateHash(newBlock) {
		return false
	}

	if newBlock.PrevHash != oldBlock.Hash {
		return false
	}

	return true
}

func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}
