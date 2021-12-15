package ipfs

import (
	"context"
	crand "crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/debridge-finance/orbitdb-go/pkg/errors"
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

// CreateRepo - create ipfs repo and apply templateCfg configs
func CreateRepo(ctx context.Context, path string, templateCfg *ipfs_cfg.Config) (ipfs_repo.Repo, error) {
	repopath := filepath.Join(path, "ipfs")
	repo, err := LoadRepoFromPath(repopath, templateCfg)
	if err != nil {
		return nil, err
	}
	rcfg, err := repo.Config()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get configs")
	}

	err = repo.SetConfig(SetIdentity(templateCfg, rcfg.Identity))
	if err != nil {
		return nil, errors.Wrap(err, "failed to set configs")
	}

	return repo, nil
}

// LoadRepoFromPath - check if repo already exisist and if not, creates new identity and applies template configs
func LoadRepoFromPath(path string, templateCfg *ipfs_cfg.Config) (ipfs_repo.Repo, error) {
	if _, err := loadPlugins(path); err != nil {
		return nil, errors.Wrap(err, "failed to load plugins")
	}
	// init repo if needed
	if !ipfs_fsrepo.IsInitialized(path) {
		identity, err := CreateIdentity()
		if err != nil {
			return nil, err
		}
		cfg := SetIdentity(templateCfg, *identity)

		if err := ipfs_fsrepo.Init(path, cfg); err != nil {
			return nil, errors.Wrap(err, "failed to init repo")
		}
	}
	return ipfs_fsrepo.Open(path)
}

// CreateTemplateConfig - creates config from file by given path (without identity)
func CreateTemplateConfig(path string) (*ipfs_cfg.Config, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfgm map[string]interface{}
	err = json.Unmarshal(content, &cfgm)
	if err != nil {
		return nil, err
	}

	cfg, err := ipfs_cfg.FromMap(cfgm)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create base config")
	}
	return cfg, nil
}

// SetIdentity - set identity
func SetIdentity(cfg *ipfs_cfg.Config, i ipfs_cfg.Identity) *ipfs_cfg.Config {
	cfg.Identity = i
	return cfg
}

// CreateIdentity - create new identity
func CreateIdentity() (*ipfs_cfg.Identity, error) {
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
	i := ipfs_cfg.Identity{}

	i.PeerID = pid.Pretty()
	i.PrivKey = base64.StdEncoding.EncodeToString(privkeyb)

	return &i, nil
}

// CreateNodeConfig - creates node config, it's differ from repo config. And we can set flags for ipfs daemon for example.
func CreateNodeConfig(repo ipfs_repo.Repo) *core.BuildCfg {
	return &core.BuildCfg{
		Online:  true,
		Routing: libp2p.DHTOption, // This option sets the node to be a full DHT node (both fetching and storing DHT Records)

		Repo:      repo,
		ExtraOpts: map[string]bool{"pubsub": true},
	}
}

// CreateIPFSNode - create IPFS node instance
func CreateIPFSNode(ctx context.Context, buildConfig *core.BuildCfg) (*core.IpfsNode, error) {
	node, err := core.NewNode(ctx, buildConfig)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// CreateCoreAPI - create IPFS coreAPI
func CreateCoreAPI(node *core.IpfsNode) (CoreAPI, error) {
	return coreapi.NewCoreAPI(node)
}

var plugins *ipfs_loader.PluginLoader

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
