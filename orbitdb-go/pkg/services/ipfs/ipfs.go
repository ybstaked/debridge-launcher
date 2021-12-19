package ipfs

import (
	"context"
	"fmt"

	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	ipfs "github.com/debridge-finance/orbitdb-go/pkg/ipfs"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
	"github.com/ipfs/go-ipfs/core"
	"github.com/libp2p/go-libp2p-core/peer"
)

type IPFS struct {
	Config Config

	log     log.Logger
	CoreAPI ipfs.CoreAPI
	Node    *core.IpfsNode
}

func Create(ctx context.Context, c Config, l log.Logger) (*IPFS, error) {

	cfg, err := ipfs.CreateTemplateConfig(c.IPFSConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create config from template")
	}

	repo, err := ipfs.CreateRepo(ctx, c.Repo, cfg)
	if err != nil {
		return nil, err
	}
	l.Info().Msg("ipfs repo was created and configurated")
	nodeCfg := ipfs.CreateNodeConfig(repo)

	l.Info().Msgf("ipfs node options was created")

	node, err := ipfs.CreateIPFSNode(ctx, nodeCfg)
	if err != nil {
		return nil, err
	}
	l.Info().Msg("ipfs node was created")

	coreapi, err := ipfs.CreateCoreAPI(node)
	if err != nil {
		return nil, err
	}
	l.Info().Msg("coreapi was created")

	return &IPFS{
		Config:  c,
		log:     l,
		CoreAPI: coreapi,
		Node:    node}, nil

}

func (i *IPFS) PeerAddrs() []string {
	peerInfo := i.Node.Peerstore.PeerInfo(peer.ID(i.Node.PeerHost.ID()))
	var peerInfoPretty []string
	for _, addr := range peerInfo.Addrs {
		addrPretty := fmt.Sprintf("%s/%s", addr, peerInfo.ID.Pretty())
		peerInfoPretty = append(peerInfoPretty, addrPretty)
	}
	return peerInfoPretty

}
