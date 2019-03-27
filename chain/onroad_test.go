package chain

import (
	"github.com/vitelabs/go-vite/common/types"
	"testing"
)

func HasOnRoadBlocks(t *testing.T, chainInstance Chain, accounts map[types.Address]*Account, addrList []types.Address) {
	for _, addr := range addrList {
		result, err := chainInstance.HasOnRoadBlocks(addr)
		if err != nil {
			t.Fatal(err)
		}
		account := accounts[addr]
		if result && len(account.UnreceivedBlocks) <= 0 {
			t.Fatal("error")
		}
		if !result && len(account.UnreceivedBlocks) > 0 {
			t.Fatal("error")
		}
	}
}

func GetOnRoadBlocksHashList(t *testing.T, chainInstance Chain, accounts map[types.Address]*Account, addrList []types.Address) {
	countPerPage := 10

	for _, addr := range addrList {
		pageNum := 0
		hashSet := make(map[types.Hash]struct{})
		account := accounts[addr]

		for {
			hashList, err := chainInstance.GetOnRoadBlocksHashList(addr, pageNum, 10)
			if err != nil {
				t.Fatal(err)
			}

			hashListLen := len(hashList)
			if hashListLen <= 0 {
				break
			}

			if hashListLen > countPerPage {
				t.Fatal(err)
			}

			for _, hash := range hashList {
				if _, ok := hashSet[hash]; ok {
					t.Fatal(err)
				}

				hashSet[hash] = struct{}{}
				hasUnReceive := false
				for _, unReceiveBlock := range account.UnreceivedBlocks {
					if unReceiveBlock.Hash == hash {
						hasUnReceive = true
						break
					}
				}

				if !hasUnReceive {
					t.Fatal("error")
				}
			}
			pageNum++
		}

		if len(hashSet) != len(account.UnreceivedBlocks) {
			t.Fatal("error")
		}
	}
}