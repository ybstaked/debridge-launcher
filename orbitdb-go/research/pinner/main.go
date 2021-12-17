package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"syscall"

	"github.com/debridge-finance/orbitdb-go/pkg/berty.tech/go-orbit-db/iface"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
	"github.com/debridge-finance/orbitdb-go/pkg/orbitdb"
	si "github.com/debridge-finance/orbitdb-go/services/ipfs"
	so "github.com/debridge-finance/orbitdb-go/services/orbitdb"
)

var (
	ErrorAlreadyExists = errors.New("address already exists")
	ErrorDoesnotExists = errors.New("address does not exists")
)

type (
	PinningList map[string]iface.Store
	Pinner      struct {
		PinningList PinningList
		odb         *so.OrbitDB
	}
)

func NewPinner(odb *so.OrbitDB) *Pinner {
	p := make(PinningList, 0)

	return &Pinner{
		PinningList: p,
		odb:         odb,
	}
}

//
func (p *Pinner) Add(ctx context.Context, address string) error {
	_, exists := p.PinningList[address]
	if exists {
		return errors.Wrap(ErrorAlreadyExists, address)
	}
	elog, err := p.odb.OrbitDB.Open(ctx, address, defaultOrbitDBOptions())
	if err != nil {
		return errors.Wrap(err, "create orbitdb eventlog")
	}
	p.PinningList[address] = elog
	return nil
}

func defaultOrbitDBOptions() *orbitdb.CreateDBOptions {
	options := &orbitdb.CreateDBOptions{}

	f := false
	options.Create = &f
	options.Overwrite = &f

	return options
}

//
func (p *Pinner) Remove(address string) error {
	_, exists := p.PinningList[address]
	if !exists {
		return errors.Wrap(ErrorDoesnotExists, address)
	}
	delete(p.PinningList, address)

	return nil
}

func (p *Pinner) replicate(ctx context.Context, store iface.Store, wg *sync.WaitGroup, errc chan error) {
	// store
	status := store.ReplicationStatus()
	fmt.Printf("status %+v\n", status)
	// evtch := replicator.Subscribe(ctx)

	wg.Done()

}

func (p *Pinner) Start(ctx context.Context, errc chan error) error {
	var wg sync.WaitGroup
	for _, store := range p.PinningList {
		wg.Add(1)
		fmt.Printf("%T\n", store)
		go p.replicate(ctx, store, &wg, errc)
	}

	wg.Wait()

	// errc <- errors.New("test error chan")
	return nil
}
func (p *Pinner) Stats(address string) error { return nil }

func main() {
	//
	rootCtx := context.Background()

	l, err := log.Create(log.DefaultConfig)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to create logger"))
	}

	ipfsCtx := context.WithValue(rootCtx, "service1", "ipfs_111")
	ipfs, err := si.Create(ipfsCtx, si.DefaultConfig, l)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to create IPFS service"))
	}
	l.Info().Msg("IPFS service was created")

	orbitdbCtx := context.WithValue(rootCtx, "service2", "orbitdbCtx_111")
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to create orbitdb service"))
	}
	odb, err := so.Create(orbitdbCtx, so.DefaultConfig, l, ipfs.CoreAPI)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to create IPFS service"))
	}
	l.Info().Msg("orbitdb service was created")

	pinner := NewPinner(odb)
	l.Info().Msg("pinner was created")

	pinnnerCtx := context.WithValue(rootCtx, "service3", "pinnerCtx_111")

	err = pinner.Add(pinnnerCtx, "/orbitdb/bafyreibx6dv6yms5225ikrsirgb2zupmc6muhk7ayoa5oekny3uef5u6yq/test")
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to add store to pinner"))
	}
	l.Info().Msg("add store to pinning list")

	errc := make(chan error)
	sig := make(chan os.Signal)

	go pinner.Start(pinnnerCtx, errc)

loop:
	for {
		shouldBrake := false
		select {
		case err := <-errc:
			if err != nil {
				l.Error().Err(err).Msg("get error from errc")
				shouldBrake = true
			}
		case si := <-sig:
			evt := l.Info().Str("signal", si.String())
			switch si {
			case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				evt = evt.Str("action", "shutdown")
				shouldBrake = true
			default:
				evt = evt.Str("action", "ignore")
			}
			evt.Msg("received signal")
		}
		if shouldBrake {
			break loop
		}
	}

}
