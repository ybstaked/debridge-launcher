package eventlog

import (
	"context"
	"encoding/json"

	"github.com/debridge-finance/orbitdb-go/pkg/berty.tech/go-ipfs-log/iface"
	"github.com/debridge-finance/orbitdb-go/pkg/berty.tech/go-orbit-db/stores/operation"
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
func (e *Eventlog) AddBinary(entry []byte) (string, error) {
	ctx := context.Background()
	// arr, err := json.Marshal(entry)
	// if err != nil {
	// 	return "", err
	// }

	h, err := e.Eventlog.Add(ctx, entry)
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
func (e *Eventlog) GetEntryOp(hash string) (operation.Operation, error) {
	ctx := context.Background()
	c, err := cid.Decode(hash)
	if err != nil {
		return nil, err
	}

	entryOp, err := e.Eventlog.Get(ctx, c)
	if err != nil {
		return nil, err
	}
	return entryOp, nil
}

type Stats struct {
	TotalEntries int32
	First        string
	Latest       string
	JSONLog      *iface.JSONLog
}

func (e *Eventlog) GetStats() *Stats {
	oplog := e.Eventlog.OpLog()
	allEntries := oplog.GetEntries()
	first := allEntries.At(uint(allEntries.Len() - 1))
	latest := allEntries.At(0)
	jsonLog := oplog.ToJSONLog()

	return &Stats{
		TotalEntries: int32(allEntries.Len()),
		JSONLog:      jsonLog,
		First:        first.GetHash().String(),
		Latest:       latest.GetHash().String(),
	}
}

func defaultOrbitDBOptions() *orbitdb.CreateDBOptions {
	options := &orbitdb.CreateDBOptions{}

	f := false
	options.Create = &f
	options.Overwrite = &f

	return options
}

func (e *Eventlog) Sync(ctx context.Context) error {

	oplog := e.Eventlog.OpLog()
	allEntries := oplog.GetEntries()
	first := allEntries.At(uint(allEntries.Len() - 1))
	next := first.GetNext()
	k, _ := allEntries.Get(next[0].KeyString())

	e.Eventlog.Sync(ctx, []iface.IPFSLogEntry{k})
	return nil
}
func (e *Eventlog) GetRoot() cid.Cid {
	oplog := e.Eventlog.OpLog()
	allEntries := oplog.GetEntries()
	first := allEntries.At(uint(allEntries.Len() - 1))
	next := first.GetNext()

	return next[0]
}

func (e *Eventlog) UnMarshalEventlogSubmissionEntry(op operation.Operation) (*EventlogSubmissionEntry, error) {

	vb := op.GetValue()
	v := &Submission{}

	err := json.Unmarshal(vb, v)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal: %v", vb)
	}
	return &EventlogSubmissionEntry{
		LogID:          op.GetEntry().GetLogID(),
		Hash:           op.GetEntry().GetHash().String(),
		AdditionalData: op.GetEntry().Copy().GetAdditionalData(),
		Next:           op.GetEntry().GetNext(),
		Refs:           op.GetEntry().GetRefs(),
		Clock:          op.GetEntry().GetClock().GetTime(),
		Value:          v,
	}, nil
}
func (e *Eventlog) UnMarshalEventlogAssetEntry(op operation.Operation) (*EventlogAssetEntry, error) {

	vb := op.GetValue()
	v := &Asset{}

	err := json.Unmarshal(vb, v)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal: %v", vb)
	}
	return &EventlogAssetEntry{
		LogID:          op.GetEntry().GetLogID(),
		Hash:           op.GetEntry().GetHash().String(),
		AdditionalData: op.GetEntry().Copy().GetAdditionalData(),
		Next:           op.GetEntry().GetNext(),
		Refs:           op.GetEntry().GetRefs(),
		Clock:          op.GetEntry().GetClock().GetTime(),
		Value:          v,
	}, nil
}

type EventlogSubmissionEntry struct {
	LogID          string            `json:"logId"`
	Hash           string            `json:"hash"`
	AdditionalData map[string]string `json:"additionalData"`
	Next           []cid.Cid         `json:"next"`
	Refs           []cid.Cid         `json:"refs"`
	Clock          int               `json:"clock"`
	Value          *Submission       `json:"entry"`
}

type Submission struct {
	SubmissionId string             `json:"submissionId"     swag_example:"f9872d1840D7322E4476C4C08c625Ab9E04d3960"`
	Signature    string             `json:"signature"`
	Payload      *SubmissionPayload `json:"payload"  swag_description:"json with payload to create new asset confirmation"`
}

type SubmissionPayload struct {
	TxHash       string `json:"txHash" swag_description:"TxHash"`
	SubmissionId string `json:"submissionId" swag_description:"SubmissionId"`
	ChainFrom    int64  `json:"chainFrom" swag_description:"ChainFrom"`
	ChainTo      int64  `json:"chainTo" swag_description:"ChainTo"`
	DebridgeId   string `json:"debridgeId" swag_description:"DebridgeId"`
	ReceiverAddr string `json:"receiverAddr" swag_description:"ReceiverAddr"`
	Amount       string `json:"amount" swag_description:"Amount"`
	EventRaw     string `json:"eventRaw" swag_description:"EventRaw"`
}

type EventlogAssetEntry struct {
	LogID          string            `json:"logId"`
	Hash           string            `json:"hash"`
	AdditionalData map[string]string `json:"additionalData"`
	Next           []cid.Cid         `json:"next"`
	Refs           []cid.Cid         `json:"refs"`
	Clock          int               `json:"clock"`
	Value          *Asset            `json:"entry"`
}

type Asset struct {
	DeployId  string        `json:"deployId"     swag_example:"f9872d1840D7322E4476C4C08c625Ab9E04d3960"`
	Signature string        `json:"signature"`
	Payload   *AssetPayload `json:"payload"  swag_description:"json with payload to create new asset confirmation"`
}
type AssetPayload struct {
	DebridgeId    string `json:"debridgeId" swag_description:"DebridgeId"`
	DeployId      string `json:"deployId" swag_description:"DeployId"`
	TokenAddress  string `json:"tokenAddress" swag_description:"TokenAddress"`
	Name          string `json:"name" swag_description:"Name"`
	Symbol        string `json:"symbol" swag_description:"Symbol"`
	NativeChainId int64  `json:"nativeChainId" swag_description:"NativeChainId"`
	Decimals      int64  `json:"decimals" swag_description:"Decimals"`
}
