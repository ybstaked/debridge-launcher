package ipfs

import (
	"context"

	ipfs "github.com/debridge-finance/orbitdb-go/pkg/ipfs"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
	"github.com/ipfs/go-ipfs/core"
)

type IPFS struct {
	Config Config

	log     log.Logger
	CoreAPI ipfs.CoreAPI
	node    *core.IpfsNode
}

func Create(ctx context.Context, c Config, l log.Logger) (*IPFS, error) {
	repo, err := ipfs.CreateRepo(ctx, c.Repo)
	if err != nil {
		return nil, err
	}
	l.Log().Msg("0 :> IPFS repo was created")
	noptions := ipfs.CreateNodeOptions(repo)
	l.Log().Msg("1 :> node options was created")

	node, err := ipfs.CreateIPFSNode(ctx, noptions)
	if err != nil {
		return nil, err
	}
	l.Log().Msg("2 :> node was created")

	coreapi, err := ipfs.CreateCoreAPI(node)
	if err != nil {
		return nil, err
	}
	l.Log().Msg("3 :> coreapi was created")

	return &IPFS{
		Config:  c,
		log:     l,
		CoreAPI: coreapi,
		node:    node}, nil

}
