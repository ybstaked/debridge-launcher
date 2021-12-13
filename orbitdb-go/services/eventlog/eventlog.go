package eventlog

import (
	"context"
	"encoding/json"

	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
	o "github.com/debridge-finance/orbitdb-go/pkg/orbitdb"
	cid "github.com/ipfs/go-cid"
)

type Eventlog struct {
	Config   Config
	Ctx      context.Context
	log      log.Logger
	Eventlog o.EventLogStore
}

func Create(ctx context.Context, c Config, l log.Logger, orbit o.OrbitDB) (*Eventlog, error) {

	elog, err := orbit.Log(ctx, "test", defaultOrbitDBOptions())
	if err != nil {
		return nil, errors.Wrap(err, "failed to create orbitdb eventlog storage")
	}

	l = l.With().Str("component", "eventlogService").Logger()
	l.Info().Msgf("eventlog storage was created: %v", elog.Address())
	all := -1
	err = elog.Load(ctx, all)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load orbitdb eventlog storage")
	}
	l.Info().Msgf("eventlog storage was loaded: %v", elog.Address())

	return &Eventlog{
		Config:   c,
		log:      l,
		Ctx:      ctx,
		Eventlog: elog,
	}, nil
}

type Entry struct {
	SubmissionId string `bson:"submissionId"`
	Signature    string `bson:"signature"`
	Event        string `bson:"event"`
}

func (e *Eventlog) Add(entry *Entry) (string, error) {
	ctx := context.Background()
	arr, err := json.Marshal(entry)
	if err != nil {
		return "", err
	}

	h, err := e.Eventlog.Add(ctx, arr)
	if err != nil {
		return "", err
	}

	return h.GetEntry().GetHash().String(), nil
}

func (e *Eventlog) Get(hash string) ([]byte, error) {
	ctx := context.Background()
	c, err := cid.Decode(hash)
	if err != nil {
		return nil, err
	}

	entryOp, err := e.Eventlog.Get(ctx, c)
	if err != nil {
		return nil, err
	}

	return entryOp.GetValue(), nil
}

type Stats struct {
	TotalOplog   int32
	TotalEntries int32
}

func (e *Eventlog) GetStats() *Stats {
	oplog := e.Eventlog.OpLog()
	allEntries := oplog.GetEntries()
	// spew.Dump([]interface{}{"oplog>>>", oplog})
	// spew.Dump([]interface{}{"allEntries>>>", allEntries})
	return &Stats{
		TotalOplog:   int32(oplog.Len()),
		TotalEntries: int32(allEntries.Len()),
	}
}

func defaultOrbitDBOptions() *o.CreateDBOptions {
	options := &o.CreateDBOptions{}

	// t := true
	f := false
	options.Create = &f
	options.Overwrite = &f

	return options
}
