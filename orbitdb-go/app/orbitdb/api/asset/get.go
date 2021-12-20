package asset

import (
	"encoding/json"

	"github.com/debridge-finance/orbitdb-go/http"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"

	"github.com/debridge-finance/orbitdb-go/pkg/services/eventlog"
)

type GetRequest struct {
	submission *eventlog.Eventlog
}

type GetRequestResult struct {
	// Hash string `json:"hash"             swag_example:"zdpuA"  swag_description:"OrbitDB hash"`
	SubmissionId string `json:"submissionId"     swag_example:"f9872d1840D7322E4476C4C08c625Ab9E04d3960"`
	Signature    string `json:"signature"`
	Event        string `json:"event"  swag_description:"json tx event with current submission"`
}

func (h *GetRequest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hash := http.PathParam(r, "hash") // FIXME: length/content control?
	res, err := h.EventlogGet(hash)
	if err != nil {
		http.WriteErrorMsg(
			w, r, http.StatusInternalServerError,
			errors.Wrapf(err, "failed to get value by hash: %v", hash),
			http.StatusText(http.StatusInternalServerError),
		)
		return
	}

	http.Write(
		w, r, http.StatusOk,
		&GetRequestResult{
			SubmissionId: res.SubmissionId,
			Signature:    res.Signature,
			Event:        res.Event,
		},
	)
}

func (h *GetRequest) EventlogGet(hash string) (*GetRequestResult, error) {
	vb, err := h.submission.Get(hash)
	if err != nil {
		return nil, err
	}

	res := &GetRequestResult{}

	err = json.Unmarshal(vb, res)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal: %v", vb)
	}

	return res, nil
}

func CreateGetRequest(e *eventlog.Eventlog) (*GetRequest, error) {
	return &GetRequest{
		submission: e,
	}, nil
}
