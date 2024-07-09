package main

import (
	"fmt"
	"math"

	"strconv"
)

const (
	blocksBucket = "ynkts"
	blockchainDB = "ynkts.db"
	targetBits   = 24
	maxNonce     = math.MaxInt64
)

func intToHex(i int64) []byte {
	return []byte(fmt.Sprintf("%x", i))
}

func main() {
	bc, err := NewBlockchain()
	if err != nil {
		fmt.Printf("%e\n", err)
		return
	}

	err = bc.AddBlock("send YNKTS to X")
	if err != nil {
		fmt.Println("add block err:", err)
		return
	}

	err = bc.AddBlock("send YNKTS to Y")
	if err != nil {
		fmt.Println("add block err:", err)
		return
	}

	bci := bc.Iterator()

	for {
		block, err := bci.Next()
		if err != nil {
			fmt.Printf("%e\n", err)
			return
		}

		pow := NewProofOfWork(block)
		fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))

		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

}
