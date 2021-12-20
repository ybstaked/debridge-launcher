package asset

import (
	"github.com/debridge-finance/orbitdb-go/http"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"

	"github.com/debridge-finance/orbitdb-go/pkg/services/eventlog"
)

type StatsRequest struct {
	eventlog *eventlog.Eventlog
}

type StatsRequestResult struct {
	// Hash string `json:"hash"             swag_example:"zdpuA"  swag_description:"OrbitDB hash"`
	TotalOplog   int32  `json:"total_oplog"     swag_example:"f9872d1840D7322E4476C4C08c625Ab9E04d3960"`
	TotalEntries int32  `json:"total_all_entries"     swag_example:"f9872d1840D7322E4476C4C08c625Ab9E04d3960"`
	FirstKey     string `json:"first"     swag_example:"f9872d1840D7322E4476C4C08c625Ab9E04d3960"`
	LastKey      string `json:"last"     swag_example:"f9872d1840D7322E4476C4C08c625Ab9E04d3960"`
}

func (h *StatsRequest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := h.EventlogStats()
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
		&StatsRequestResult{
			TotalOplog:   res.TotalOplog,
			TotalEntries: res.TotalEntries,
			FirstKey:     res.FirstKey,
			LastKey:      res.LastKey,
		},
	)
}

func (h *StatsRequest) EventlogStats() (*StatsRequestResult, error) {
	total := h.eventlog.GetStats()

	return &StatsRequestResult{
		TotalOplog:   int32(total.TotalOplog),
		TotalEntries: int32(total.TotalEntries),
		FirstKey:     total.FirstKey,
		LastKey:      total.LastKey,
	}, nil
}

func CreateStatsRequest(e *eventlog.Eventlog) (*StatsRequest, error) {
	return &StatsRequest{
		eventlog: e,
	}, nil
}
