package contracts

import (
	"github.com/vitelabs/go-vite/common/helper"
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/go-vite/vm/abi"
	"math/big"
	"strings"
	"testing"
)

func TestContractsABIInit(t *testing.T) {
	tests := []string{jsonRegister, jsonVote, jsonPledge, jsonConsensusGroup, jsonMintage}
	for _, data := range tests {
		if _, err := abi.JSONToABIContract(strings.NewReader(jsonRegister)); err != nil {
			t.Fatalf("json to abi failed, %v, %v", data, err)
		}
	}
}

func BenchmarkRegisterUnpackVariable(b *testing.B) {
	value := helper.HexToBytes("0000000000000000000000000000000000000000000000000000000000000100000000000000000000000000988dd19d15702dbf8a4d316f920b1fdcf57d4a50000000000000000000000000988dd19d15702dbf8a4d316f920b1fdcf57d4a50000000000000000000000000988dd19d15702dbf8a4d316f920b1fdcf57d4a500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005c18dd820000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000066e6f646532350000000000000000000000000000000000000000000000000000")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		registration := new(Registration)
		ABIRegister.UnpackVariable(registration, VariableNameRegistration, value)
	}
}

func BenchmarkVoteUnpackVariable(b *testing.B) {
	value := helper.HexToBytes("0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000a73757065724e6f64653100000000000000000000000000000000000000000000")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vote := new(VoteInfo)
		ABIVote.UnpackVariable(vote, VariableNameVoteStatus, value)
	}
}

func BenchmarkConsensusGroupUnpackVariable(b *testing.B) {
	value := helper.HexToBytes("0000000000000000000000000000000000000000000000000000000000000019000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000220000000000000000000000000b3db179e6ae63aa8d2114c386ebe91c9aa470ab50000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005ba78e1600000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000d3c21bcecceda10000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000076a7000000000000000000000000000000000000000000000000000000000000000000")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		info := new(ConsensusGroupInfo)
		ABIConsensusGroup.UnpackVariable(info, VariableNameConsensusGroupInfo, value)
	}
}

func TestPackMethodParam(t *testing.T) {
	_, err := PackMethodParam(AddressVote, MethodNameVote, types.DELEGATE_GID, "node")
	if err != nil {
		t.Fatalf("pack method param failed, %v", err)
	}
}

func TestPackConsensusGroupConditionParam(t *testing.T) {
	_, err := PackConsensusGroupConditionParam(RegisterConditionPrefix, uint8(1), big.NewInt(1), ledger.ViteTokenId, uint64(10))
	if err != nil {
		t.Fatalf("pack consensus group condition param failed")
	}
}