package pinner

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/debridge-finance/orbitdb-go/pkg/log"

	// config "github.com/ipfs/go-ipfs-config"
	config "github.com/ipfs/go-ipfs-config"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/core/node/libp2p"
	"github.com/ipfs/go-ipfs/plugin/loader" // This package is needed so that all the preloaded plugins are loaded automatically
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	icore "github.com/ipfs/interface-go-ipfs-core"
)

type (
	CoreAPI = icore.CoreAPI
)

// /// ------ Setting up the IPFS Repo
func createRepo(path string) (string, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			return "", err
		}
	}
	cfg, err := config.Init(ioutil.Discard, 2048)
	if err != nil {
		return "", err
	}
	// When creating the repository, you can define custom settings on the repository, such as enabling experimental
	// features (See experimental-features.md) or customizing the gateway endpoint.
	// To do such things, you should modify the variable `cfg`. For example:

	// https://github.com/ipfs/go-ipfs/blob/master/docs/experimental-features.md#ipfs-filestore
	cfg.Experimental.FilestoreEnabled = true
	// https://github.com/ipfs/go-ipfs/blob/master/docs/experimental-features.md#ipfs-urlstore
	cfg.Experimental.UrlstoreEnabled = true
	// https://github.com/ipfs/go-ipfs/blob/master/docs/experimental-features.md#directory-sharding--hamt
	cfg.Experimental.ShardingEnabled = true
	// https://github.com/ipfs/go-ipfs/blob/master/docs/experimental-features.md#ipfs-p2p
	cfg.Experimental.Libp2pStreamMounting = true
	// https://github.com/ipfs/go-ipfs/blob/master/docs/experimental-features.md#p2p-http-proxy
	cfg.Experimental.P2pHttpProxy = true
	// https://github.com/ipfs/go-ipfs/blob/master/docs/experimental-features.md#strategic-providing
	cfg.Experimental.StrategicProviding = true

	fmt.Println("repoPath", path)
	// Create the repo with the config
	err = fsrepo.Init(path, cfg)
	if err != nil {
		return "", fmt.Errorf("failed to init ephemeral node: %s", err)
	}

	return path, nil
}

func setupPlugins(externalPluginsPath string) error {
	// Load any external plugins if available on externalPluginsPath
	plugins, err := loader.NewPluginLoader(filepath.Join(externalPluginsPath, "plugins"))
	if err != nil {
		return fmt.Errorf("error loading plugins: %s", err)
	}

	// Load preloaded and external plugins
	if err := plugins.Initialize(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	if err := plugins.Inject(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	return nil
}

// Creates an IPFS node and returns its coreAPI
func newCoreAPI(ctx context.Context, repoPath string) (CoreAPI, error) {
	// Open the repo
	repo, err := fsrepo.Open(repoPath)
	if err != nil {
		return nil, err
	}

	// Construct the node

	nodeOptions := &core.BuildCfg{
		Online:  true,
		Routing: libp2p.DHTOption, // This option sets the node to be a full DHT node (both fetching and storing DHT Records)
		// Routing:   libp2p.DHTClientOption, // This option sets the node to be a client DHT node (only fetching records)
		Repo:      repo,
		ExtraOpts: map[string]bool{"pubsub": true},
	}

	node, err := core.NewNode(ctx, nodeOptions)
	if err != nil {
		return nil, err
	}
	fmt.Printf("node.Identity %v\n", node.Identity)
	fmt.Printf("node.Repo %v\n", node.Repo)
	fmt.Printf("node %v\n", node)

	// Attach the Core API to the constructed node
	return coreapi.NewCoreAPI(node)
}

func Create(ctx context.Context, l log.Logger, path string) (icore.CoreAPI, error) {
	if err := setupPlugins(""); err != nil {
		return nil, err
	}
	repoPath, err := createRepo(path)
	if err != nil {
		fmt.Printf("failed to create repo: %v\n", err)
		return nil, err
	}
	return newCoreAPI(ctx, repoPath)

}
