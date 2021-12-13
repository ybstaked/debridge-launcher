package ipfs

import (
	"context"
	crand "crypto/rand"
	"encoding/base64"
	"path/filepath"

	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	ipfs_ds "github.com/ipfs/go-datastore"
	ipfs_cfg "github.com/ipfs/go-ipfs-config"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/core/node/libp2p" // This package is needed so that all the preloaded plugins are loaded automatically
	ipfs_loader "github.com/ipfs/go-ipfs/plugin/loader"
	ipfs_repo "github.com/ipfs/go-ipfs/repo"
	ipfs_fsrepo "github.com/ipfs/go-ipfs/repo/fsrepo"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
	p2p_ci "github.com/libp2p/go-libp2p-core/crypto"
	p2p_peer "github.com/libp2p/go-libp2p-core/peer"
)

type (
	CoreAPI = coreiface.CoreAPI
)

func CreateRepo(ctx context.Context, path string) (ipfs_repo.Repo, error) {
	repopath := filepath.Join(path, "ipfs")
	repo, err := LoadRepoFromPath(repopath)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func CreateNodeOptions(repo ipfs_repo.Repo) *core.BuildCfg {
	return &core.BuildCfg{
		Online:  true,
		Routing: libp2p.DHTOption, // This option sets the node to be a full DHT node (both fetching and storing DHT Records)

		Repo:      repo,
		ExtraOpts: map[string]bool{"pubsub": true},
	}
}

func CreateIPFSNode(ctx context.Context, nodeOptions *core.BuildCfg) (*core.IpfsNode, error) {
	// Construct the node
	node, err := core.NewNode(ctx, nodeOptions)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func CreateCoreAPI(node *core.IpfsNode) (CoreAPI, error) {
	return coreapi.NewCoreAPI(node)
}

// defaultConnMgrHighWater is the default value for the connection managers
// 'high water' mark
const defaultConnMgrHighWater = 500

// defaultConnMgrLowWater is the default value for the connection managers 'low
// water' mark
const defaultConnMgrLowWater = 400

// defaultConnMgrGracePeriod is the default value for the connection managers
// grace period
const defaultConnMgrGracePeriod = time.Second * 20

// @NOTE(gfanton): this will be removed with gomobile-ipfs
var plugins *ipfs_loader.PluginLoader

func CreateMockedRepo(dstore ipfs_ds.Batching) (ipfs_repo.Repo, error) {
	c, err := createBaseConfig()
	if err != nil {
		return nil, err
	}

	return &ipfs_repo.Mock{
		D: dstore,
		C: *c,
	}, nil
}

func LoadRepoFromPath(path string) (ipfs_repo.Repo, error) {
	if _, err := loadPlugins(path); err != nil {
		return nil, errors.Wrap(err, "failed to load plugins")
	}

	// init repo if needed
	if !ipfs_fsrepo.IsInitialized(path) {
		cfg, err := createBaseConfig()
		if err != nil {
			return nil, errors.Wrap(err, "failed to create base config")
		}

		ucfg, err := upgradeToPersistanceConfig(cfg)
		if err != nil {
			return nil, errors.Wrap(err, "failed to upgrade repo")
		}

		if err := ipfs_fsrepo.Init(path, ucfg); err != nil {
			return nil, errors.Wrap(err, "failed to init repo")
		}
	}

	return ipfs_fsrepo.Open(path)
}

// save
var DefaultSwarmListeners = []string{
	"/ip4/0.0.0.0/tcp/0",
	"/ip6/::/tcp/0",
	"/ip4/0.0.0.0/udp/0/quic",
	"/ip6/::/udp/0/quic",
}

// save
func createBaseConfig() (*ipfs_cfg.Config, error) {
	priv, pub, err := p2p_ci.GenerateKeyPairWithReader(p2p_ci.Ed25519, 2048, crand.Reader) // nolint:staticcheck
	if err != nil {
		return nil, err
	}

	pid, err := p2p_peer.IDFromPublicKey(pub) // nolint:staticcheck
	if err != nil {
		return nil, err
	}

	privkeyb, err := p2p_ci.MarshalPrivateKey(priv)
	if err != nil {
		return nil, err
	}

	c := ipfs_cfg.Config{}
	spew.Dump([]interface{}{">>> ipfs config swarm", c.Swarm})

	// set default bootstrap
	c.Bootstrap = ipfs_cfg.DefaultBootstrapAddresses
	c.Peering.Peers = []p2p_peer.AddrInfo{}

	// Identity
	c.Identity.PeerID = pid.Pretty()
	c.Identity.PrivKey = base64.StdEncoding.EncodeToString(privkeyb)

	// Discovery
	c.Discovery.MDNS.Enabled = true
	c.Discovery.MDNS.Interval = 20

	// swarm listeners
	c.Addresses.Swarm = DefaultSwarmListeners

	// c.Swarm
	// Swarm
	// c.Swarm.AutoRelay.Enabled = true
	// c.Swarm.EnableRelayHop = false
	c.Swarm.ConnMgr = ipfs_cfg.ConnMgr{
		LowWater:    defaultConnMgrLowWater,
		HighWater:   defaultConnMgrHighWater,
		GracePeriod: defaultConnMgrGracePeriod.String(),
		Type:        "basic",
	}

	c.Routing = ipfs_cfg.Routing{
		Type: "dhtclient",
	}

	return &c, nil
}

// save
func upgradeToPersistanceConfig(cfg *ipfs_cfg.Config) (*ipfs_cfg.Config, error) {
	cfgCopy, err := cfg.Clone()
	if err != nil {
		return nil, err
	}

	// setup the node mount points.
	cfgCopy.Mounts = ipfs_cfg.Mounts{
		IPFS: "/ipfs",
		IPNS: "/ipns",
	}

	cfgCopy.Ipns = ipfs_cfg.Ipns{
		ResolveCacheSize: 128,
	}

	cfgCopy.Reprovider = ipfs_cfg.Reprovider{
		Interval: "12h",
		Strategy: "all",
	}

	cfgCopy.Datastore = ipfs_cfg.Datastore{
		StorageMax:         "10GB",
		StorageGCWatermark: 90, // 90%
		GCPeriod:           "1h",
		BloomFilterSize:    0,
		Spec: map[string]interface{}{
			"type": "mount",
			"mounts": []interface{}{
				map[string]interface{}{
					"mountpoint": "/blocks",
					"type":       "measure",
					"prefix":     "flatfs.datastore",
					"child": map[string]interface{}{
						"type":      "flatfs",
						"path":      "blocks",
						"sync":      true,
						"shardFunc": "/repo/flatfs/shard/v1/next-to-last/2",
					},
				},
				map[string]interface{}{
					"mountpoint": "/",
					"type":       "measure",
					"prefix":     "leveldb.datastore",
					"child": map[string]interface{}{
						"type":        "levelds",
						"path":        "datastore",
						"compression": "none",
					},
				},
			},
		},
	}

	return cfgCopy, nil
}

// save
func loadPlugins(repoPath string) (*ipfs_loader.PluginLoader, error) {
	if plugins != nil {
		return plugins, nil
	}

	pluginpath := filepath.Join(repoPath, "plugins")

	lp, err := ipfs_loader.NewPluginLoader(pluginpath)
	if err != nil {
		return nil, err
	}

	if err = lp.Initialize(); err != nil {
		return nil, err
	}

	if err = lp.Inject(); err != nil {
		return nil, err
	}

	plugins = lp
	return lp, nil
}
