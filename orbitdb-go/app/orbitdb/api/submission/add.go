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
	SubmissionId string `json:"submissionId"     swag_example:"f9872d1840D7322E4476C4C08c625Ab9E04d3960"`
	Signature    string `json:"signature"`
	Event        string `json:"event"  swag_description:"json tx event with current submission"`
}

type AddRequestResult struct {
	Hash string `json:"hash"             swag_example:"zdpuA"  swag_description:"OrbitDB hash"`
}

func (h *AddRequest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		params = &AddRequestParams{}
		e      = &eventlog.Entry{}
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

	err = h.ParametersToModel(params, e)
	if err != nil {
		http.WriteErrorMsg(
			w, r, http.StatusInternalServerError,
			errors.Wrap(err, "failed to cast parameters to emitent model"),
			http.StatusText(http.StatusInternalServerError),
		)
		return
	}

	hash, err := h.AddToEventlog(e)
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

func (h *AddRequest) AddToEventlog(e *eventlog.Entry) (string, error) {

	hash, err := h.eventlog.Add(e)
	if err != nil {
		return "", err
	}
	h.log.Debug().Msg("AddToOrbitdb")

	return hash, nil
}
func (h *AddRequest) ParametersToModel(p *AddRequestParams, e *eventlog.Entry) error {
	e.SubmissionId = p.SubmissionId
	e.Signature = p.Signature
	e.Event = p.Event

	return nil
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
