package api

import (
	"github.com/vitelabs/go-vite/chain"
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/log15"
	apidefi "github.com/vitelabs/go-vite/rpcapi/api/defi"
	"github.com/vitelabs/go-vite/rpcapi/api/dex"
	"github.com/vitelabs/go-vite/vite"
	"github.com/vitelabs/go-vite/vm/contracts/defi"
)

type DeFiApi struct {
	vite  *vite.Vite
	chain chain.Chain
	log   log15.Logger
}

func NewDeFiApi(vite *vite.Vite) *DeFiApi {
	return &DeFiApi{
		vite:  vite,
		chain: vite.Chain(),
		log:   log15.New("module", "rpc_api/defi_api"),
	}
}

func (f DeFiApi) String() string {
	return "DeFiApi"
}

type RpcBaseAccount struct {
	Available   string `json:"available"`
	Subscribing string `json:"subscribing,omitempty"`
	Subscribed  string `json:"subscribed,omitempty"`
	Invested    string `json:"invested,omitempty"`
	Locked      string `json:"locked,omitempty"`
}

type RpcLoanAccount struct {
	Available string `json:"available"`
	Invested  string `json:"invested"`
}

type DeFiAccount struct {
	Token       *RpcTokenInfo   `json:"token"`
	BaseAccount *RpcBaseAccount `json:"baseAccount,omitempty"`
	LoanAccount *RpcLoanAccount `json:"loanAccount,omitempty"`
}

func (f DeFiApi) GetAccountInfo(addr types.Address, tokenId *types.TokenTypeId) (map[types.TokenTypeId]*DeFiAccount, error) {
	db, err := getVmDb(f.chain, types.AddressDeFi)
	if err != nil {
		return nil, err
	}
	fund, _ := defi.GetFund(db, addr)
	defiAccs, ok := defi.GetAccounts(fund, tokenId)
	if !ok {
		return nil, nil
	}
	accounts := make(map[types.TokenTypeId]*DeFiAccount, 0)
	for _, v := range defiAccs {
		token, _ := types.BytesToTokenTypeId(v.Token)
		tokenInfo, err := f.chain.GetTokenInfoById(token)
		if err != nil {
			return nil, err
		}
		deFiAcc := &DeFiAccount{}
		deFiAcc.Token = RawTokenInfoToRpc(tokenInfo, token)
		if v.BaseAccount != nil {
			baseAccount := &RpcBaseAccount{}
			if v.BaseAccount.Available != nil {
				baseAccount.Available = dex.AmountBytesToString(v.BaseAccount.Available)
			}
			if v.BaseAccount.Subscribing != nil {
				baseAccount.Subscribing = dex.AmountBytesToString(v.BaseAccount.Subscribing)
			}
			if v.BaseAccount.Subscribed != nil {
				baseAccount.Subscribed = dex.AmountBytesToString(v.BaseAccount.Subscribed)
			}
			if v.BaseAccount.Invested != nil {
				baseAccount.Invested = dex.AmountBytesToString(v.BaseAccount.Invested)
			}
			if v.BaseAccount.Locked != nil {
				baseAccount.Locked = dex.AmountBytesToString(v.BaseAccount.Locked)
			}
			deFiAcc.BaseAccount = baseAccount
		}
		if v.LoanAccount != nil {
			loanAccount := &RpcLoanAccount{}
			if v.LoanAccount.Available != nil {
				loanAccount.Available = dex.AmountBytesToString(v.LoanAccount.Available)
			}
			if v.LoanAccount.Invested != nil {
				loanAccount.Invested = dex.AmountBytesToString(v.LoanAccount.Invested)
			}
			deFiAcc.LoanAccount = loanAccount
		}
		accounts[token] = deFiAcc
	}
	return accounts, nil
}

func (f DeFiApi) GetLoanInfo(loanId uint64) (*apidefi.RpcLoan, error) {
	db, err := getVmDb(f.chain, types.AddressDeFi)
	if err != nil {
		return nil, err
	}
	if loan, ok := defi.GetLoan(db, loanId); ok {
		return apidefi.LoanToRpc(loan), nil
	} else {
		return nil, defi.LoanNotExistsErr
	}
}

func (f DeFiApi) GetSubscriptionInfo(loanId uint64, address types.Address) (*apidefi.RpcSubscription, error) {
	db, err := getVmDb(f.chain, types.AddressDeFi)
	if err != nil {
		return nil, err
	}
	if subscription, ok := defi.GetSubscription(db, loanId, address.Bytes()); ok {
		return apidefi.SubscriptionToRpc(subscription), nil
	} else {
		return nil, defi.SubscriptionNotExistsErr
	}
}

func (f DeFiApi) GetInvest(investId uint64) (*apidefi.RpcInvest, error) {
	db, err := getVmDb(f.chain, types.AddressDeFi)
	if err != nil {
		return nil, err
	}
	if invest, ok := defi.GetInvest(db, investId); ok {
		return apidefi.InvestToRpc(invest), nil
	} else {
		return nil, defi.InvestNotExistsErr
	}
}

func (f DeFiApi) GetSbpRegistration(investId uint64) (*apidefi.RpcSBPRegistration, error) {
	db, err := getVmDb(f.chain, types.AddressDeFi)
	if err != nil {
		return nil, err
	}
	if invest, ok := defi.GetInvest(db, investId); ok {
		if sbpReg, ok := defi.GetSBPRegistration(db, invest.InvestHash); ok {
			return apidefi.SBPRegistrationToRpc(sbpReg), nil
		} else {
			return nil, defi.SBPRegistrationNotExistsErr
		}
	} else {
		return nil, defi.InvestNotExistsErr
	}
}

func (f DeFiApi) GetInvestQuotaInfo(investId uint64) (*apidefi.RpcInvestQuotaInfo, error) {
	db, err := getVmDb(f.chain, types.AddressDeFi)
	if err != nil {
		return nil, err
	}
	if invest, ok := defi.GetInvest(db, investId); ok {
		if investQuota, ok := defi.GetInvestQuotaInfo(db, invest.InvestHash); ok {
			return apidefi.InvestQuotaInfoToRpc(investQuota), nil
		} else {
			return nil, defi.InvalidQuotaInvestErr
		}
	} else {
		return nil, defi.InvestNotExistsErr
	}
}

func (f DeFiApi) GetTimestamp() (int64, error) {
	db, err := getVmDb(f.chain, types.AddressDeFi)
	if err != nil {
		return 0, err
	}
	return defi.GetDeFiTimestamp(db), nil
}