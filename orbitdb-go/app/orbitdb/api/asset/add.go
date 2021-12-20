package asset

import (
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
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
	DeployId  string   `json:"deployId"     swag_example:"f9872d1840D7322E4476C4C08c625Ab9E04d3960"`
	Signature string   `json:"signature"`
	Payload   *Payload `json:"payload"  swag_description:"json with payload to create new asset confirmation"`
}

type Payload struct {
	DebridgeId    string `json:"debridgeId" swag_description:"DebridgeId"`
	DeployId      string `json:"deployId" swag_description:"DeployId"`
	TokenAddress  string `json:"tokenAddress" swag_description:"TokenAddress"`
	Name          string `json:"name" swag_description:"Name"`
	Symbol        string `json:"symbol" swag_description:"Symbol"`
	NativeChainId int64  `json:"nativeChainId" swag_description:"NativeChainId"`
	Decimals      int64  `json:"decimals" swag_description:"Decimals"`
}

type AddRequestResult struct {
	Hash string `json:"hash"             swag_example:"zdpuA"  swag_description:"OrbitDB hash"`
}

func (h *AddRequest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		params = &AddRequestParams{}
		// payload     = &eventlog.Entry{}
	)

	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		http.WriteError(w, r, http.StatusBadRequest, err)
		return
	}
	r.Body.Close()

	err = params.Validate()
	if err != nil {
		http.WriteError(
			w, r, http.StatusBadRequest,
			errors.Wrap(err, "failed to validate request parameters"),
		)
		return
	}
	spew.Dump([]interface{}{"params", params})
	// err = h.ParametersToModel(params, e)
	// if err != nil {
	// 	http.WriteErrorMsg(
	// 		w, r, http.StatusInternalServerError,
	// 		errors.Wrap(err, "failed to cast parameters to emitent model"),
	// 		http.StatusText(http.StatusInternalServerError),
	// 	)
	// 	return
	// }

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
	h.log.Debug().Msg("AddToEventlogAsset")

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
	if p.Payload == nil {
		return errors.New("payload field should not be empty")
	}
	return nil
}

func CreateAddRequest(c Config, l log.Logger, e *eventlog.Eventlog) (*AddRequest, error) {
	return &AddRequest{
		Config:   c,
		log:      l,
		eventlog: e,
	}, nil
}
