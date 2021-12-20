package pinner

import (
	"context"

	"github.com/debridge-finance/orbitdb-go/pkg/log"
	"github.com/debridge-finance/orbitdb-go/pkg/orbitdb"
	// "github.com/debridge-finance/orbitdb-go/pkg/pinner"
)

type Pinner struct {
	ctx context.Context
	log log.Logger

	Config Config
	// Pinner *pinner.Pinner
}

func Create(ctx context.Context, c Config, l log.Logger, odb *orbitdb.OrbitDB) (*Pinner, error) {
	l = l.With().Str("component", "pinner-service").Logger()

	// p := pinner.New(odb)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to create new pinner")
	// }
	l.Info().Msg("pinner service was created")

	return &Pinner{
		ctx:    ctx,
		log:    l,
		Config: c,
		// Pinner: p,
	}, nil
}

// type Entry struct {
// 	SubmissionId string `bson:"submissionId"`
// 	Signature    string `bson:"signature"`
// 	Event        string `bson:"event"`
// }

// func (e *Eventlog) Add(entry *Entry) (string, error) {
// 	ctx := context.Background()
// 	arr, err := json.Marshal(entry)
// 	if err != nil {
// 		return "", err
// 	}

// 	h, err := e.Eventlog.Add(ctx, arr)
// 	if err != nil {
// 		return "", err
// 	}

// 	return h.GetEntry().GetHash().String(), nil
// }

// func (e *Eventlog) Get(hash string) ([]byte, error) {
// 	ctx := context.Background()
// 	c, err := cid.Decode(hash)
// 	if err != nil {
// 		return nil, err
// 	}

// 	entryOp, err := e.Eventlog.Get(ctx, c)
// 	if err != nil {
// 		return nil, err
// 	}
// 	spew.Dump([]interface{}{"operation", entryOp.GetOperation()})

// 	return entryOp.GetValue(), nil
// }

// type Stats struct {
// 	TotalOplog   int32
// 	TotalEntries int32
// 	FirstKey     string
// 	LastKey      string
// }

// func (e *Eventlog) GetStats() *Stats {
// 	oplog := e.Eventlog.OpLog()
// 	allEntries := oplog.GetEntries()
// 	keys := allEntries.Keys()
// 	fKey, lKey := "", ""
// 	if len(keys) > 0 {
// 		fKey = keys[0]
// 		lKey = keys[len(keys)-1]
// 	}

// 	return &Stats{
// 		TotalOplog:   int32(oplog.Len()),
// 		TotalEntries: int32(allEntries.Len()),
// 		FirstKey:     fKey,
// 		LastKey:      lKey,
// 	}
// }

// func defaultOrbitDBOptions() *o.CreateDBOptions {
// 	options := &o.CreateDBOptions{}

// 	f := false
// 	options.Create = &f
// 	options.Overwrite = &f

// 	return options
// }
