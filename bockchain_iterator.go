package main

import "github.com/boltdb/bolt"

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (i *BlockchainIterator) Next() (*Block, error) {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		bc, err := Deserialize(encodedBlock)
		if err != nil {
			return err
		}

		block = bc
		return nil
	})
	if err != nil {
		return nil, err
	}

	i.currentHash = block.PrevBlockHash
	return block, nil
}
