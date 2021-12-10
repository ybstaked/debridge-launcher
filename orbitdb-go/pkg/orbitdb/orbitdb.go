package orbitdb

import (
	"context"
	"fmt"

	"os"

	odb "github.com/debridge-finance/orbitdb-go/pkg/berty.tech/go-orbit-db"
	i "github.com/debridge-finance/orbitdb-go/pkg/ipfs"
)

type (
	OrbitDB = odb.OrbitDB
)

func Create(ctx context.Context, ipfs i.CoreAPI, dbPath string) (*OrbitDB, error) {
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
