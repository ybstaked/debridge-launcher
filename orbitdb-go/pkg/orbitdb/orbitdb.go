package orbitdb

import (
	"context"
	"fmt"

	"os"

	odb "berty.tech/go-orbit-db"
	coreapi "github.com/ipfs/interface-go-ipfs-core"
)

type (
	OrbitDB = odb.OrbitDB
)

func Create(ctx context.Context, ipfs coreapi.CoreAPI, dbPath string) (*OrbitDB, error) {
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(dbPath, 0755)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}

	options := &odb.NewOrbitDBOptions{Directory: &dbPath}

	orbitdb, err := odb.NewOrbitDB(ctx, ipfs, options)
	if err != nil {
		fmt.Printf("failed to create NewOrbitDB: %v", err)
	}
	return &orbitdb, nil
}
