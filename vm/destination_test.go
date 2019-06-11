package vm

import (
	"bytes"
	"encoding/hex"
	"github.com/vitelabs/go-vite/common/types"
	"math/big"
	"testing"
)

func TestCodeAnalysis(t *testing.T) {
	tests := []struct {
		code   []byte
		result []byte
	}{
		{[]byte{byte(PUSH1), 0x01, 0x01, 0x01}, []byte{0x40, 0x00, 0x00, 0x00, 0x00}},
		{[]byte{byte(PUSH1), byte(PUSH1), byte(PUSH1), byte(PUSH1)}, []byte{0x50, 0x00, 0x00, 0x00, 0x00}},
		{[]byte{byte(PUSH8), byte(PUSH8), byte(PUSH8), byte(PUSH8), byte(PUSH8), byte(PUSH8), byte(PUSH8), byte(PUSH8), 0x01, 0x01, 0x01}, []byte{0x7f, 0x80, 0x00, 0x00, 0x00, 0x00}},
		{[]byte{byte(PUSH8), 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}, []byte{0x7f, 0x80, 0x00, 0x00, 0x00, 0x00}},
		{[]byte{0x01, 0x01, 0x01, 0x01, 0x01, byte(PUSH2), byte(PUSH2), byte(PUSH2), 0x01, 0x01, 0x01}, []byte{0x03, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{[]byte{0x01, 0x01, 0x01, 0x01, 0x01, byte(PUSH2), 0x01, 0x01, 0x01, 0x01, 0x01}, []byte{0x03, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{[]byte{byte(PUSH3), 0x01, 0x01, 0x01, byte(PUSH1), 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}, []byte{0x74, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{[]byte{byte(PUSH3), 0x01, 0x01, 0x01, byte(PUSH1), 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}, []byte{0x74, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{[]byte{0x01, byte(PUSH8), 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}, []byte{0x3f, 0xc0, 0x00, 0x00, 0x00, 0x00}},
		{[]byte{0x01, byte(PUSH8), 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}, []byte{0x3f, 0xc0, 0x00, 0x00, 0x00, 0x00}},
		{[]byte{byte(PUSH16), 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}, []byte{0x7f, 0xff, 0x80, 0x00, 0x00, 0x00}},
		{[]byte{byte(PUSH16), 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}, []byte{0x7f, 0xff, 0x80, 0x00, 0x00, 0x00}},
		{[]byte{byte(PUSH16), 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}, []byte{0x7f, 0xff, 0x80, 0x00, 0x00, 0x00}},
		{[]byte{byte(PUSH8), 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, byte(PUSH1), 0x01}, []byte{0x7f, 0xa0, 0x00, 0x00, 0x00, 0x00}},
		{[]byte{byte(PUSH8), 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, byte(PUSH1), 0x01}, []byte{0x7f, 0xa0, 0x00, 0x00, 0x00, 0x00}},
		{[]byte{byte(PUSH32)}, []byte{0x7f, 0xff, 0xff, 0xff, 0x80}},
		{[]byte{byte(PUSH32)}, []byte{0x7f, 0xff, 0xff, 0xff, 0x80}},
		{[]byte{byte(PUSH32)}, []byte{0x7f, 0xff, 0xff, 0xff, 0x80}},
	}
	for _, test := range tests {
		ret := codeBitmap(test.code)
		if !bytes.Equal(test.result, ret) {
			t.Fatalf("analysis fail, got %v, expected %v", ret, test.result)
		}
	}
}

func TestHas(t *testing.T) {
	tests := []struct {
		code   []byte
		dest   *big.Int
		result bool
	}{
		{[]byte{byte(PUSH1), byte(JUMPDEST)}, big.NewInt(1), false},
		{[]byte{byte(PUSH1), 0, byte(JUMPDEST)}, big.NewInt(2), true},
		{[]byte{byte(PUSH1), 0, byte(JUMPDEST)}, big.NewInt(1), false},
		{[]byte{byte(PUSH32), 0, byte(JUMPDEST)}, big.NewInt(2), false},
	}
	for _, test := range tests {
		d := make(destinations)
		result := d.has(types.Address{}, test.code, test.dest)
		if result != test.result {
			t.Fatalf("analysis result error, code: [%v], dest: %v, expected: %v, got: %v", test.code, test.dest, test.result, result)
		}
	}
}

func TestContainsAuxCode(t *testing.T) {
	code := "608060405260043610604c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632d16c91a146052578063ba3d93d614609f57604c565b60006000fd5b348015605e5760006000fd5b5060896004803603602081101560745760006000fd5b810190808035906020019092919050505060c8565b6040518082815260200191505060405180910390f35b34801560ab5760006000fd5b5060b260dd565b6040518082815260200191505060405180910390f35b60008160006000505401905060d8565b919050565b6000600060005054905060eb565b9056fea165627a7a72305820614b83ac7ad21936973699f9740d3cdc61d467686393dccfdda77dc053f21b7c0029"
	codeB, _ := hex.DecodeString(code)
	if !containsAuxCode(codeB) {
		t.Fatalf("check contains aux code failed")
	}
}

func TestContainsStatusCode2(t *testing.T) {
	code := []byte{byte(HEIGHT), byte(PUSH1), 0}
	if !ContainsStatusCode(code) {
		t.Fatalf("check contains status code failed")
	}
	code = []byte{byte(PUSH1), byte(HEIGHT)}
	if ContainsStatusCode(code) {
		t.Fatalf("check contains status code failed")
	}
}

func TestGetCodeWithoutAuxCodeAndParams(t *testing.T) {
	testCases := []struct {
		name       string
		code       string
		resultCode string
	}{
		/*{
			"normal",
			"608060405234801561001057600080fd5b50610141806100206000396000f3fe608060405260043610610041576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806391a6cb4b14610046575b600080fd5b6100896004803603602081101561005c57600080fd5b81019080803574ffffffffffffffffffffffffffffffffffffffffff16906020019092919050505061008b565b005b8074ffffffffffffffffffffffffffffffffffffffffff164669ffffffffffffffffffff163460405160405180820390838587f1505050508074ffffffffffffffffffffffffffffffffffffffffff167faa65281f5df4b4bd3c71f2ba25905b907205fce0809a816ef8e04b4d496a85bb346040518082815260200191505060405180910390a25056fea165627a7a7230582017ef8aba505f3b1f219622dd9c54e93003fb280d0d7c48364215f545da9575b20029",
			"608060405234801561001057600080fd5b50610141806100206000396000f3fe608060405260043610610041576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806391a6cb4b14610046575b600080fd5b6100896004803603602081101561005c57600080fd5b81019080803574ffffffffffffffffffffffffffffffffffffffffff16906020019092919050505061008b565b005b8074ffffffffffffffffffffffffffffffffffffffffff164669ffffffffffffffffffff163460405160405180820390838587f1505050508074ffffffffffffffffffffffffffffffffffffffffff167faa65281f5df4b4bd3c71f2ba25905b907205fce0809a816ef8e04b4d496a85bb346040518082815260200191505060405180910390a25056fe",
		},*/
		{
			"exception",
			"608060405234801561001057600080fd5b50610141806100206000396000f3fe608069fea165627a7a7230582060405260043610610041576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806391a6cb4b14610046575b600080fd5b6100896004803603602081101561005c57600080fd5b81019080803574ffffffffffffffffffffffffffffffffffffffffff16906020019092919050505061008b565b005b8074ffffffffffffffffffffffffffffffffffffffffff164669ffffffffffffffffffff163460405160405180820390838587f1505050508074ffffffffffffffffffffffffffffffffffffffffff167faa65281f5df4b4bd3c71f2ba25905b907205fce0809a816ef8e04b4d496a85bb346040518082815260200191505060405180910390a25056fea165627a7a7230582017ef8aba505f3b1f219622dd9c54e93003fb280d0d7c48364215f545da9575b20029",
			"608060405234801561001057600080fd5b50610141806100206000396000f3fe608069fea165627a7a7230582060405260043610610041576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806391a6cb4b14610046575b600080fd5b6100896004803603602081101561005c57600080fd5b81019080803574ffffffffffffffffffffffffffffffffffffffffff16906020019092919050505061008b565b005b8074ffffffffffffffffffffffffffffffffffffffffff164669ffffffffffffffffffff163460405160405180820390838587f1505050508074ffffffffffffffffffffffffffffffffffffffffff167faa65281f5df4b4bd3c71f2ba25905b907205fce0809a816ef8e04b4d496a85bb346040518082815260200191505060405180910390a25056fe",
		},
	}
	for _, testCase := range testCases {
		codeB, _ := hex.DecodeString(testCase.code)
		resultCode := getCodeWithoutAuxCodeAndParams(codeB)
		resultCodeStr := hex.EncodeToString(resultCode)
		if resultCodeStr != testCase.resultCode {
			t.Fatalf("get code without auxCode and params failed,\n expectCode: %v,\n resultCode: %v", testCase.resultCode, resultCodeStr)
		}
	}

}
