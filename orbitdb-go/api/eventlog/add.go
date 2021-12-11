package eventlog

import (
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
	"github.com/debridge-finance/orbitdb-go/http"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"

	orbit "github.com/debridge-finance/orbitdb-go/services/orbitdb"
)

type AddRequest struct {
	Config Config

	log     log.Logger
	orbitdb *orbit.OrbitDB // TODO: change to eventlog
}

type RequestParams struct {
	SubmissionId string `json:"submissionId"     swag_example:"f9872d1840D7322E4476C4C08c625Ab9E04d3960"`
	Signature    string `json:"signature"`
	Event        string `json:"event"  swag_description:"json tx event with current submission"`
}

type Entry struct {
	SubmissionId string `bson:"submissionId"`
	Signature    string `bson:"signature"`
	Event        string `bson:"event"`
}

type RequestResult struct {
	Hash string `json:"hash"             swag_example:"zdpuA"  swag_description:"OrbitDB hash"`
}

func (h *AddRequest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		params = &RequestParams{}
		e      = &Entry{}
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

	hash, err := h.AddToOrbitdb(e)
	if err != nil {
		http.WriteErrorMsg(
			w, r, http.StatusInternalServerError,
			errors.Wrap(err, "failed to subscribe to txs in notification service"),
			http.StatusText(http.StatusInternalServerError),
		)
		return
	}
	http.Write(
		w, r, http.StatusOk,
		&RequestResult{
			Hash: hash,
		},
	)
}

func (h *AddRequest) AddToOrbitdb(e *Entry) (string, error) {
	// identity := emitent.Identity(*p.Identity)

	// e.Id = rand.String(32)
	// e.Identity = emitent.IdentityContainer{
	// 	Current:  &identity,
	// 	Versions: emitent.IdentityVersions{},
	// }
	// e.Address = p.Address
	h.log.Debug().Msg("AddToOrbitdb")

	return "ipfs-hash-whould-be-here", nil
}
func (h *AddRequest) ParametersToModel(p *RequestParams, e *Entry) error {
	// identity := emitent.Identity(*p.Identity)
	spew.Dump([]interface{}{"request params", p})
	// e.Id = rand.String(32)
	// e.Identity = emitent.IdentityContainer{
	// 	Current:  &identity,
	// 	Versions: emitent.IdentityVersions{},
	// }
	// e.Address = p.Address

	return nil
}

//

func (p *RequestParams) SetDefaults() {
loop:
	for {
		switch {
		default:
			break loop
		}
	}
}

func (p *RequestParams) Validate() error {
	// var err error

	return nil
}

////

//

// func CreateAddRequest() (*AddRequest, error) {
// 	return &AddRequest{
// 		orbitdb: nil,
// 	}, nil
// }

func CreateAddRequest(c Config, l log.Logger, odb *orbit.OrbitDB) (*AddRequest, error) {
	return &AddRequest{
		Config:  c,
		log:     l,
		orbitdb: odb,
	}, nil
}
