package main

import (
	"bytes"
	"encoding/gob"
)

func Serialize(b *Block) ([]byte, error) {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}

func Deserialize(d []byte) (*Block, error) {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		return nil, err
	}

	return &block, nil
}
