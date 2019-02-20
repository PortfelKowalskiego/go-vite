package nodemanager

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vitelabs/go-vite/cmd/utils"
	"github.com/vitelabs/go-vite/node"
	"gopkg.in/urfave/cli.v1"
)

type RecoverNodeManager struct {
	ctx  *cli.Context
	node *node.Node
}

func NewRecoverNodeManager(ctx *cli.Context, maker NodeMaker) (*RecoverNodeManager, error) {
	node, err := maker.MakeNode(ctx)
	if err != nil {
		return nil, err
	}

	// single mode
	node.Config().Single = true
	node.ViteConfig().Net.Single = true

	return &RecoverNodeManager{
		ctx:  ctx,
		node: node,
	}, nil
}

func (nodeManager *RecoverNodeManager) getDeleteToHeight() uint64 {
	deleteToHeight := uint64(0)
	if nodeManager.ctx.GlobalIsSet(utils.LedgerDeleteToHeight.Name) {
		deleteToHeight = nodeManager.ctx.GlobalUint64(utils.LedgerDeleteToHeight.Name)
	}
	return deleteToHeight
}

func (nodeManager *RecoverNodeManager) Start() error {
	// Start up the node
	node := nodeManager.node
	err := StartNode(nodeManager.node)
	if err != nil {
		return err
	}

	deleteToHeight := nodeManager.getDeleteToHeight()
	c := node.Vite().Chain()

	if deleteToHeight <= 0 {
		err := errors.New("deleteToHeight is 0.\n")
		panic(err)
	}

	fmt.Printf("Deleting to %d...\n", deleteToHeight)

	if _, _, err := c.DeleteSnapshotBlocksToHeight(deleteToHeight); err != nil {
		fmt.Printf("Delete to %d height failed. error is "+err.Error()+"\n", deleteToHeight)
		return err
	}
	fmt.Printf("Delete to %d successed!\n", deleteToHeight)

	fmt.Printf("Rebuild data...\n")
	if err := c.TrieGc().Recover(); err != nil {
		fmt.Errorf("Rebuild data failed! error is %s\n", err.Error())
	} else {
		fmt.Printf("Rebuild data successed!\n")
	}

	fmt.Printf("Latest snapshot block height is %d\n", c.GetLatestSnapshotBlock().Height)
	return nil
}

func (nodeManager *RecoverNodeManager) Stop() error {

	StopNode(nodeManager.node)

	return nil
}

func (nodeManager *RecoverNodeManager) Node() *node.Node {
	return nodeManager.node
}
