package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"syscall"
	"time"

	"github.com/debridge-finance/orbitdb-go/pkg/berty.tech/go-orbit-db/iface"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
	"github.com/debridge-finance/orbitdb-go/pkg/orbitdb"
	si "github.com/debridge-finance/orbitdb-go/pkg/services/ipfs"
	so "github.com/debridge-finance/orbitdb-go/pkg/services/orbitdb"
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
		lock        sync.RWMutex
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
	p.lock.Lock()
	defer p.lock.Unlock()

	_, exists := p.PinningList[address]
	if exists {
		return errors.Wrap(ErrorAlreadyExists, address)
	}
	elog, err := p.odb.OrbitDB.Log(ctx, address, defaultOrbitDBOptions())
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

func (p *Pinner) replicate(ctx context.Context, store iface.Store, wg *sync.WaitGroup, errc chan error) {
	defer wg.Done()

	for {
		fmt.Printf("%v\tprogress: %v\tbufferLen: %v\n", store.Address(), store.ReplicationStatus().GetProgress(), store.Replicator().GetBufferLen())
		// v, ok := <-ch
		// if !ok {
		// 	break

		// }
		// fmt.Printf("queue: %v\n", store.Replicator().GetQueue())
		// fmt.Printf("value: %v\n", v)
		// entries := store.OpLog().GetEntries().Slice()
		// var first, _ log_iface.IPFSLogEntry
		// if len(entries) > 0 {
		// 	first = entries[0]
		// 	// last = entries[len(entries)-1]
		// }
		// first.GetHash()

		// fmt.Printf("%v\tprogress: %v\tbufferLen: %v\t\first:\ntime:%+v\thash:%v\nlast:\ntime:%+v\thash:%v\n", store.Address(), store.ReplicationStatus().GetProgress(), store.Replicator().GetBufferLen(), first.GetClock(), first.GetHash().Hash(), last.GetClock(), last.GetHash().Hash())

		time.Sleep(10 * time.Second)

	}
	// store
	// evtch := replicator.Subscribe(ctx)

}

type PinnerStats struct {
}

func (p *Pinner) Stats(address string) error {

	return nil
}

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
	s1 := "/orbitdb/bafyreihrgrfr4m74wkesroranlra7yf2kgqs2n3icfy76flwzzngjw3sba/test" // london
	s2 := "/orbitdb/bafyreigzsnrsa62hpw6udbi7nmxhhdgi46d5wxfpoiqa2lgm2ovu6tkewu/test" // frankfurt
	err = pinner.Add(pinnnerCtx, s1)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to add store to pinner"))
	}
	l.Info().Msgf("add store %v to pinning list", s1)
	err = pinner.Add(pinnnerCtx, s2)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to add store to pinner"))
	}
	l.Info().Msgf("add store %v to pinning list", s2)

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
