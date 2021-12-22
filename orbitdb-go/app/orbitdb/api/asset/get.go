package asset

import (
	"github.com/debridge-finance/orbitdb-go/http"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"

	"github.com/debridge-finance/orbitdb-go/pkg/services/eventlog"
)

type GetRequest struct {
	asset *eventlog.Eventlog
}

type GetRequestResult struct {
	DeployId  string   `json:"deployId"     swag_example:"f9872d1840D7322E4476C4C08c625Ab9E04d3960"`
	Signature string   `json:"signature"`
	Payload   *Payload `json:"payload"  swag_description:"json with payload to create new asset confirmation"`
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
		res,
	)
}

type GetEntryRequestResult = eventlog.EventlogAssetEntry

func (h *GetRequest) EventlogGet(hash string) (*GetEntryRequestResult, error) {
	op, err := h.asset.GetEntryOp(hash)
	if err != nil {
		return nil, err
	}

	entry, err := h.asset.UnMarshalEventlogAssetEntry(op)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal entry: %v", op)
	}

	return entry, nil
}

func CreateGetRequest(e *eventlog.Eventlog) (*GetRequest, error) {
	return &GetRequest{
		asset: e,
	}, nil
}
