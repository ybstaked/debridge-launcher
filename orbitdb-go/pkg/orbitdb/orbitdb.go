package pinner

import (
	"context"
	"fmt"

	"os"

	orbitdb "berty.tech/go-orbit-db"
	coreapi "github.com/ipfs/interface-go-ipfs-core"
)

func CreateOrbitdb(ctx context.Context, ipfs coreapi.CoreAPI, dbPath string) (orbitdb.OrbitDB, error) {
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(dbPath, 0755)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}

	options := &orbitdb.NewOrbitDBOptions{Directory: &dbPath}

	orbitdb, err := orbitdb.NewOrbitDB(ctx, ipfs, options)
	if err != nil {
		fmt.Printf("failed to create NewOrbitDB: %v", err)
	}
	return orbitdb, nil
}
