package chainmgr

import (
	"github.com/ontio/ontology/account"
	"math"
	"github.com/ontio/ontology/core/types"
)

func GetChainManager() *ChainManager {
	return defaultChainManager
}

func GetAccount() *account.Account {
	chainmgr := GetChainManager()
	return chainmgr.account
}

func GetParentShardID() uint64 {
	chainmgr := GetChainManager()
	return chainmgr.parentShardID
}

func GetParentBlockHeight() uint64 {
	chainmgr := GetChainManager()
	chainmgr.lock.RLock()
	defer chainmgr.lock.RUnlock()

	m := chainmgr.blockPool.Shards[chainmgr.parentShardID]
	if m == nil {
		return math.MaxUint64
	}

	h := uint64(0)
	for _, blk := range m {
		if blk.BlockHeight > h {
			h = blk.BlockHeight
		}
	}

	return h
}

func GetParentBlockHeader(height uint64) *types.Header {
	chainmgr := GetChainManager()
	chainmgr.lock.RLock()
	defer chainmgr.lock.RUnlock()

	m := chainmgr.blockPool.Shards[chainmgr.parentShardID]
	if m == nil {
		return nil
	}
	if blk, present := m[height]; present && blk != nil {
		return blk.Header.Header
	}

	return nil
}
