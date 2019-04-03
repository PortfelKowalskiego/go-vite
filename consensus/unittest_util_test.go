package consensus

import (
	"os"
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
)

var GenesisJson = "{\"GenesisAccountAddress\":\"vite_ab24ef68b84e642c0ddca06beec81c9acb1977bbd7da27a87a\",\"ForkPoints\":{},\"ConsensusGroupInfo\":{\"ConsensusGroupInfoMap\":{\"00000000000000000001\":{\"NodeCount\":2,\"Interval\":1,\"PerCount\":3,\"RandCount\":2,\"RandRank\":100,\"Repeat\":1,\"CheckLevel\":0,\"CountingTokenId\":\"tti_5649544520544f4b454e6e40\",\"RegisterConditionId\":1,\"RegisterConditionParam\":{\"PledgeAmount\":100000000000000000000000,\"PledgeHeight\":1,\"PledgeToken\":\"tti_5649544520544f4b454e6e40\"},\"VoteConditionId\":1,\"VoteConditionParam\":{},\"Owner\":\"vite_ab24ef68b84e642c0ddca06beec81c9acb1977bbd7da27a87a\",\"PledgeAmount\":0,\"WithdrawHeight\":1},\"00000000000000000002\":{\"NodeCount\":2,\"Interval\":3,\"PerCount\":1,\"RandCount\":2,\"RandRank\":100,\"Repeat\":48,\"CheckLevel\":1,\"CountingTokenId\":\"tti_5649544520544f4b454e6e40\",\"RegisterConditionId\":1,\"RegisterConditionParam\":{\"PledgeAmount\":100000000000000000000000,\"PledgeHeight\":1,\"PledgeToken\":\"tti_5649544520544f4b454e6e40\"},\"VoteConditionId\":1,\"VoteConditionParam\":{},\"Owner\":\"vite_ab24ef68b84e642c0ddca06beec81c9acb1977bbd7da27a87a\",\"PledgeAmount\":0,\"WithdrawHeight\":1}},\"RegistrationInfoMap\":{\"00000000000000000001\":{\"s1\":{\"NodeAddr\":\"vite_360232b0378111b122685a15e612143dc9a89cfa7e803f4b5a\",\"PledgeAddr\":\"vite_360232b0378111b122685a15e612143dc9a89cfa7e803f4b5a\",\"Amount\":100000000000000000000000,\"WithdrawHeight\":7776000,\"RewardTime\":1,\"CancelTime\":0,\"HisAddrList\":[\"vite_360232b0378111b122685a15e612143dc9a89cfa7e803f4b5a\"]},\"s2\":{\"NodeAddr\":\"vite_ce18b99b46c70c8e6bf34177d0c5db956a8c3ea7040a1c1e25\",\"PledgeAddr\":\"vite_ce18b99b46c70c8e6bf34177d0c5db956a8c3ea7040a1c1e25\",\"Amount\":100000000000000000000000,\"WithdrawHeight\":7776000,\"RewardTime\":1,\"CancelTime\":0,\"HisAddrList\":[\"vite_ce18b99b46c70c8e6bf34177d0c5db956a8c3ea7040a1c1e25\"]}}}},\"MintageInfo\":{\"TokenInfoMap\":{\"tti_5649544520544f4b454e6e40\":{\"TokenName\":\"Vite Token\",\"TokenSymbol\":\"VITE\",\"TotalSupply\":1000000000000000000000000000,\"Decimals\":18,\"Owner\":\"vite_ab24ef68b84e642c0ddca06beec81c9acb1977bbd7da27a87a\",\"PledgeAmount\":0,\"PledgeAddr\":\"vite_ab24ef68b84e642c0ddca06beec81c9acb1977bbd7da27a87a\",\"WithdrawHeight\":0,\"MaxSupply\":115792089237316195423570985008687907853269984665640564039457584007913129639935,\"OwnerBurnOnly\":false,\"IsReIssuable\":true}},\"LogList\":[{\"Data\":\"\",\"Topics\":[\"3f9dcc00d5e929040142c3fb2b67a3be1b0e91e98dac18d5bc2b7817a4cfecb6\",\"000000000000000000000000000000000000000000005649544520544f4b454e\"]}]},\"PledgeInfo\":{\"PledgeBeneficialMap\":{\"vite_360232b0378111b122685a15e612143dc9a89cfa7e803f4b5a\":1000000000000000000000,\"vite_ce18b99b46c70c8e6bf34177d0c5db956a8c3ea7040a1c1e25\":1000000000000000000000}},\"AccountBalanceMap\":{\"vite_ab24ef68b84e642c0ddca06beec81c9acb1977bbd7da27a87a\":{\"tti_5649544520544f4b454e6e40\":899999000000000000000000000},\"vite_56fd05b23ff26cd7b0a40957fb77bde60c9fd6ebc35f809c23\":{\"tti_5649544520544f4b454e6e40\":100000000000000000000000000},\"vite_360232b0378111b122685a15e612143dc9a89cfa7e803f4b5a\":{\"tti_5649544520544f4b454e6e40\":899999000000000000000000000},\"vite_ce18b99b46c70c8e6bf34177d0c5db956a8c3ea7040a1c1e25\":{\"tti_5649544520544f4b454e6e40\":100000000000000000000000000}}}"

var UnitTestDir = "testdata-unittest"

func NewDb(t *testing.T, dirName string) *leveldb.DB {
	db, err := leveldb.OpenFile(dirName, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	return db
}

func ClearDb(t *testing.T, dirName string) {
	os.RemoveAll(dirName)
}
