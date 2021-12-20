package eventlog

import (
	"context"
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
	"github.com/debridge-finance/orbitdb-go/pkg/orbitdb"
	cid "github.com/ipfs/go-cid"
)

type Eventlog struct {
	Config   Config
	Ctx      context.Context
	log      log.Logger
	Eventlog orbitdb.EventLogStore
}

func Create(ctx context.Context, c Config, l log.Logger, odb orbitdb.OrbitDB) (*Eventlog, error) {

	elog, err := odb.Log(ctx, c.Name, defaultOrbitDBOptions())
	if err != nil {
		return nil, errors.Wrap(err, "failed to create orbitdb eventlog storage")
	}

	l = l.With().Str("component", "eventlogService").Logger()
	l.Info().Msgf("eventlog storage was created: %v", elog.Address())
	err = elog.Load(ctx, c.Limit)
	if err != nil {
		l.Error().Msgf("eventlog storage was loaded: %v", elog.Address())
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
	spew.Dump([]interface{}{"operation", entryOp.GetOperation()})

	return entryOp.GetValue(), nil
}

type Stats struct {
	TotalOplog   int32
	TotalEntries int32
	FirstKey     string
	LastKey      string
}

func (e *Eventlog) GetStats() *Stats {
	oplog := e.Eventlog.OpLog()
	allEntries := oplog.GetEntries()
	keys := allEntries.Keys()
	fKey, lKey := "", ""
	if len(keys) > 0 {
		fKey = keys[0]
		lKey = keys[len(keys)-1]
	}

	return &Stats{
		TotalOplog:   int32(oplog.Len()),
		TotalEntries: int32(allEntries.Len()),
		FirstKey:     fKey,
		LastKey:      lKey,
	}
}

func defaultOrbitDBOptions() *orbitdb.CreateDBOptions {
	options := &orbitdb.CreateDBOptions{}

	f := false
	options.Create = &f
	options.Overwrite = &f

	return options
}
