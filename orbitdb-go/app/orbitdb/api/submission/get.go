package submission

import (
	"github.com/debridge-finance/orbitdb-go/http"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"

	"github.com/debridge-finance/orbitdb-go/pkg/services/eventlog"
)

type GetEntryRequest struct {
	submission *eventlog.Eventlog
}

func (h *GetEntryRequest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

type GetEntryRequestResult = eventlog.EventlogSubmissionEntry

func (h *GetEntryRequest) EventlogGet(hash string) (*GetEntryRequestResult, error) {
	op, err := h.submission.GetEntryOp(hash)
	if err != nil {
		return nil, err
	}

	entry, err := h.submission.UnMarshalEventlogSubmissionEntry(op)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal entry: %v", op)
	}

	return entry, nil
}

func CreateGetEntryRequest(e *eventlog.Eventlog) (*GetEntryRequest, error) {
	return &GetEntryRequest{
		submission: e,
	}, nil
}
