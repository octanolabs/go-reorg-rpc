package chain

import (
	"time"

	"github.com/octanolabs/go-reorg-rpc/common"

	lru "github.com/hashicorp/golang-lru"
	log "github.com/sirupsen/logrus"
)

const (
	blockCacheLimit = 100
)

// Fake chain
type Chain struct {
	blockCache *lru.Cache
	head       uint64
	reorg      bool
}

// Returns a new fake Chain with blockCache limit 'cacheLimit'
func New(cacheLimit int) *Chain {
	log.Println("Generating new chain")
	bCache, _ := lru.New(cacheLimit)
	return &Chain{bCache, 0, true}
}

func (c *Chain) InitGenesis() {
	c.addBlock(0, genesis())
}

func (c *Chain) GetHead() string {
	return common.EncodeUint64(c.head)
}

func (c *Chain) GetBlock(n uint64) common.Block {
	var b common.Block
	if cached, ok := c.blockCache.Get(n); ok {
		b = cached.(common.Block)
		return b
	} else {
		log.Warnf("GetBlock not found: %v", n)
		return b
	}
}

func (c *Chain) StartMiner(blockTime string) {
	interval, err := time.ParseDuration(blockTime)
	if err != nil {
		log.Fatalf("Can't parse blockTime: %v", err)
	}
	ticker := time.NewTicker(interval)
	log.Printf("Mining new blocks with interval: %v", interval)

	go func() {
		for {
			select {
			case <-ticker.C:
				var height = c.head + 1
				parent, _ := c.blockCache.Get(c.head)
				b := genBlock(height, parent.(common.Block).Hash)
				c.addBlock(height, b)
				log.Printf("Mined new block, height: %v, hash: %v", height, b.Hash)
				if c.head%20 == 0 {
					if c.reorg {
						log.Warnln("Faking reorg, depth: 3")
						c.head = c.head - 3
						c.reorg = false
					} else {
						c.reorg = true
					}
				} else {
					c.head = height
				}
			}
		}
	}()
}

func genBlock(height uint64, parentHash string) common.Block {
	var b common.Block
	var txns []string
	var uncles []string

	b.Number = common.EncodeUint64(height)
	b.Timestamp = common.EncodeUint64(uint64(time.Now().Unix()))
	b.Transactions = txns
	b.Hash = randomHash(64)
	b.ParentHash = parentHash
	b.Sha3Uncles = randomHash(64)
	b.Miner = randomHash(40)
	b.Difficulty = "0x12a05f2000"
	b.TotalDifficulty = "0x2540be4000"
	b.Size = "0x21b"
	b.GasUsed = "0x0"
	b.GasLimit = "0x8000000"
	b.Nonce = randomHash(16)
	b.Uncles = uncles
	b.MixHash = randomHash(64)
	b.ReceiptsRoot = randomHash(64)
	b.StateRoot = randomHash(64)
	b.ExtraData = "0x"
	b.TransactionRoot = randomHash(64)
	b.LogsBloom = "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	return b
}

func (c *Chain) addBlock(height uint64, b common.Block) {
	c.blockCache.Add(height, b)
	c.head = height
}
