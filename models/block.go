package models

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

// Block represents each 'item' in the blockchain
type Block struct {
	ID        uint   `gorm:"primary_key"`
	Index     int    `json:"index"`
	Timestamp string `json:"timestamp"`
	BS        string `json:"bs" gorm:"COMMENT:'业务数据';size:65535"`
	Hash      string `json:"hash"`
	PrevHash  string `json:"prev_hash"`
}

// create a new block using previous block's hash
func GenerateBlock(oldBlock Block, BS string) Block {
	var newBlock Block
	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BS = BS
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

// make sure block is valid by checking index, and comparing the hash of the previous block
func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func creationBlock() Block {
	t := time.Now()
	genesisBlock := Block{}
	genesisBlock = Block{
		Index:     0,
		Timestamp: t.String(),
		BS:        "",
		Hash:      calculateHash(genesisBlock),
		PrevHash:  ""}
	//spew.Dump(genesisBlock)
	return genesisBlock
}

// SHA256 hasing
func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.BS + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func ReadBlocks() ([]Block, error) {
	var blocks []Block
	err := db.Raw("select * from block order by block.index").Scan(&blocks).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		block := creationBlock()
		blocks = append(blocks, block)
	}
	if len(blocks) > 0 {
		return blocks, nil
	} else {
		block := creationBlock()
		blocks = append(blocks, block)
		return blocks, nil
	}
}

func WriteBlocks(blocks []Block) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}
	for _, block := range blocks {
		err := tx.Where("block.index=?", block.Index).Save(&block).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
