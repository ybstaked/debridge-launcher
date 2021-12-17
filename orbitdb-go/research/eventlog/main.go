package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/debridge-finance/orbitdb-go/pkg/berty.tech/go-orbit-db/events"
	"github.com/debridge-finance/orbitdb-go/pkg/berty.tech/go-orbit-db/stores/replicator"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
	"github.com/debridge-finance/orbitdb-go/pkg/orbitdb"
	si "github.com/debridge-finance/orbitdb-go/services/ipfs"
	so "github.com/debridge-finance/orbitdb-go/services/orbitdb"
)

var address = "/orbitdb/bafyreihrgrfr4m74wkesroranlra7yf2kgqs2n3icfy76flwzzngjw3sba/test"

func main() {
	//
	rootCtx := context.Background()

	l, err := log.Create(log.DefaultConfig)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to create logger"))
	}

	// IPFS
	ipfsCtx := context.WithValue(rootCtx, "service1", "ipfs_111")
	ipfs, err := si.Create(ipfsCtx, si.DefaultConfig, l)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to create IPFS service"))
	}
	// l.Info().Msgf("IPFS service was created\n%v", ipfs.PeerAddrs())

	orbitdbCtx := context.WithValue(rootCtx, "service2", "orbitdbCtx_111")
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to create orbitdb service"))
	}
	odb, err := so.Create(orbitdbCtx, so.DefaultConfig, l, ipfs.CoreAPI)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to create IPFS service"))
	}
	l.Info().Msg("orbitdb service was created")

	// Eventlog

	eventlogCtx := context.WithValue(rootCtx, "servicer", "eventlog_111")
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to create orbitdb service"))
	}
	logStore, err := odb.Log(eventlogCtx, address, defaultOrbitDBOptions())
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to create logStore"))
	}
	l.Info().Msg("logStore service was created")
	// err = logStore.Load(eventlogCtx, 200)
	// if err != nil {
	// 	fmt.Println(errors.Wrap(err, "failed to load logStore"))
	// }
	// l.Info().Msgf("logStore service was loaded (200). Length: %v", logStore.OpLog().GetEntries().Len())

	// logStore.Sync()

	replicator := logStore.Replicator()

	ch := replicator.Subscribe(eventlogCtx)
	var wg sync.WaitGroup
	wg.Add(1)
	go listenReplicates(ch, &wg, replicator)
	wg.Wait()
	spew.Dump([]interface{}{"1entries amount: ", logStore.OpLog().GetEntries().Len()})
	time.Sleep(10 * time.Second)
	spew.Dump([]interface{}{"2entries amount: ", logStore.OpLog().GetEntries().Len()})
	time.Sleep(10 * time.Second)
	spew.Dump([]interface{}{"3entries amount: ", logStore.OpLog().GetEntries().Len()})

	// for _, e := range es {
	// 	l.Info().Msgf("entry: %v", e)
	// }

	// logStoreCh := make(chan operation.Operation)
	// all := -1
	// err = logStore.Stream(eventlogCtx, logStoreCh, &iface.StreamOptions{Amount: &all})
	// if err != nil {
	// 	fmt.Println(errors.Wrap(err, "failed to create logStore stream channel"))
	// }
	// go func() {
	// 	l.Info().Msgf("get value from logStoreCh: %v\n", <-logStoreCh)
	// }()

	// op, err := logStore.Add(eventlogCtx, []byte("example-value"))
	// if err != nil {
	// 	fmt.Println(errors.Wrap(err, "failed to add value to the logStore"))
	// }

	// l.Info().Msgf("get value from logStoreCh: %v\n", op.GetEntry().GetHash().String())

}

func listenReplicates(ch <-chan events.Event, wg *sync.WaitGroup, r replicator.Replicator) {
	defer wg.Done()
	i := 200
	for {
		v := <-ch
		fmt.Printf("bufferLen: %v\n", r.GetBufferLen())
		fmt.Printf("queue: %v\n", r.GetQueue())
		fmt.Printf("value: %v\n", v)
		i--
		if i <= 0 {
			break
		}

	}
}

func defaultOrbitDBOptions() *orbitdb.CreateDBOptions {
	options := &orbitdb.CreateDBOptions{}

	f := false
	options.Create = &f
	options.Overwrite = &f

	return options
}
