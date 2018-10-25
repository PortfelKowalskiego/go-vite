package chain

import (
	"github.com/vitelabs/go-vite/common"
	"github.com/vitelabs/go-vite/config"
	"path/filepath"
)

var innerChainInstance Chain

func getChainInstance() Chain {
	if innerChainInstance == nil {
		//home := common.HomeDir()

		innerChainInstance = NewChain(&config.Config{
			//DataDir: filepath.Join(common.HomeDir(), "govite_testdata"),

			DataDir: filepath.Join(common.HomeDir(), "Library/GVite/testdata"),
			//Chain: &config.Chain{
			//	KafkaProducers: []*config.KafkaProducer{{
			//		Topic:      "test",
			//		BrokerList: []string{"abc", "def"},
			//	}},
			//},
		})
		innerChainInstance.Init()
		innerChainInstance.Start()
	}

	return innerChainInstance
}