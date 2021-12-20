package submission

import (
	"encoding/json"

	"github.com/debridge-finance/orbitdb-go/http"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"

	"github.com/debridge-finance/orbitdb-go/pkg/services/eventlog"
)

type AddRequest struct {
	Config Config

	log      log.Logger
	eventlog *eventlog.Eventlog // TODO: change to eventlog
}

type AddRequestParams struct {
	SubmissionId string   `json:"submissionId"     swag_example:"f9872d1840D7322E4476C4C08c625Ab9E04d3960"`
	Signature    string   `json:"signature"`
	Payload      *Payload `json:"payload"  swag_description:"json with payload to create new submission"`
}

type Payload struct {
	TxHash       string `json:"txHash" swag_description:"TxHash"`
	SubmissionId string `json:"submissionId" swag_description:"SubmissionId"`
	ChainFrom    int64  `json:"chainFrom" swag_description:"ChainFrom"`
	ChainTo      int64  `json:"chainTo" swag_description:"ChainTo"`
	DebridgeId   string `json:"debridgeId" swag_description:"DebridgeId"`
	ReceiverAddr string `json:"receiverAddr" swag_description:"ReceiverAddr"`
	Amount       string `json:"amount" swag_description:"Amount"`
	EventRaw     string `json:"eventRaw" swag_description:"EventRaw"`
}

type AddRequestResult struct {
	Hash string `json:"hash"             swag_example:"zdpuA"  swag_description:"OrbitDB hash"`
}

func (h *AddRequest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		params = &AddRequestParams{}
	)

	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		http.WriteError(w, r, http.StatusBadRequest, err)
		return
	}
	r.Body.Close()

	params.SetDefaults()
	err = params.Validate()
	if err != nil {
		http.WriteError(
			w, r, http.StatusBadRequest,
			errors.Wrap(err, "failed to validate request parameters"),
		)
		return
	}

	hash, err := h.AddToEventlog(params)
	if err != nil {
		http.WriteErrorMsg(
			w, r, http.StatusInternalServerError,
			errors.Wrap(err, "failed to add entry to the eventlog"),
			http.StatusText(http.StatusInternalServerError),
		)
		return
	}
	http.Write(
		w, r, http.StatusOk,
		&AddRequestResult{
			Hash: hash,
		},
	)
}

func (h *AddRequest) AddToEventlog(p *AddRequestParams) (string, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	hash, err := h.eventlog.AddBinary(b)
	if err != nil {
		return "", err
	}
	h.log.Debug().Msg("AddToEventlogSubmission")

	return hash, nil
}

func (p *AddRequestParams) SetDefaults() {
loop:
	for {
		switch {
		default:
			break loop
		}
	}
}

func (p *AddRequestParams) Validate() error {
	// var err error

	return nil
}

func CreateAddRequest(c Config, l log.Logger, e *eventlog.Eventlog) (*AddRequest, error) {
	return &AddRequest{
		Config:   c,
		log:      l,
		eventlog: e,
	}, nil
}
