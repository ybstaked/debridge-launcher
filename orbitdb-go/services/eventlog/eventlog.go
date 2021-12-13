package eventlog

import (
	"context"
	"encoding/json"

	"github.com/debridge-finance/orbitdb-go/pkg/berty.tech/go-orbit-db/stores/replicator"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
	o "github.com/debridge-finance/orbitdb-go/pkg/orbitdb"
	cid "github.com/ipfs/go-cid"
)

type Eventlog struct {
	Config Config

	log      log.Logger
	Eventlog o.EventLogStore
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

func (e *Eventlog) GetStats() replicator.ReplicationInfo {
	return e.Eventlog.ReplicationStatus()
}

func Create(ctx context.Context, c Config, l log.Logger, orbit o.OrbitDB) (*Eventlog, error) {
	elog, err := orbit.Log(ctx, "test", nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create orbitdb eventlog storage")
	}

	l = l.With().Str("component", "orbitdbService").Logger()

	return &Eventlog{
		Config:   c,
		log:      l,
		Eventlog: elog,
	}, nil
}
