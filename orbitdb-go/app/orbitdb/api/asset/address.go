package asset

import (
	"github.com/debridge-finance/orbitdb-go/http"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"

	"github.com/debridge-finance/orbitdb-go/pkg/services/eventlog"
)

type AddressRequest struct {
	eventlog *eventlog.Eventlog
}

type AddressRequestResult struct {
	OrbitLogsDb string `json:"OrbitLogsDb"     swag_example:"/orbitdb/bafyreihatxlrreu2axly2klwweb4f33d3qjwyvx6opbflci464pgyk22o4/test"`
}

func (h *AddressRequest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := h.EventlogAddress()
	if err != nil {
		http.WriteErrorMsg(
			w, r, http.StatusInternalServerError,
			errors.Wrap(err, "failed to get eventlog stats"),
			http.StatusText(http.StatusInternalServerError),
		)
		return
	}

	http.Write(
		w, r, http.StatusOk,
		res,
	)
}

func (h *AddressRequest) EventlogAddress() (*AddressRequestResult, error) {
	a := h.eventlog.Eventlog.Address()

	return &AddressRequestResult{
		OrbitLogsDb: a.String(),
	}, nil
}

func CreateAddressRequest(e *eventlog.Eventlog) (*AddressRequest, error) {
	return &AddressRequest{
		eventlog: e,
	}, nil
}
