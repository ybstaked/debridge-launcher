package submission

import (
	"github.com/debridge-finance/orbitdb-go/http"
	"github.com/debridge-finance/orbitdb-go/pkg/berty.tech/go-ipfs-log/iface"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"

	"github.com/debridge-finance/orbitdb-go/pkg/services/eventlog"
)

type StatsRequest struct {
	eventlog *eventlog.Eventlog
}

type StatsRequestResult struct {
	TotalEntries int32          `json:"total_all_entries"     swag_example:"f9872d1840D7322E4476C4C08c625Ab9E04d3960"`
	First        string         `json:"first"     swag_example:"bafyreibdlu2wo5gdxacru2tphlsqht43pullbesmt6tpbs3w35vjze4ja4"`
	Latest       string         `json:"latest"     swag_example:"bafyreifdabdk5oozhsotwypzu24nchcioj5u3rioohxsbqzn6vw7yitbfi"`
	JSONLog      *iface.JSONLog `json:"jsonLog"     swag_example:"bafyreifdabdk5oozhsotwypzu24nchcioj5u3rioohxsbqzn6vw7yitbfi"`
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
			TotalEntries: res.TotalEntries,
			First:        res.First,
			Latest:       res.Latest,
			JSONLog:      res.JSONLog,
		},
	)
}

// func

func (h *StatsRequest) EventlogStats() (*StatsRequestResult, error) {
	total := h.eventlog.GetStats()

	return &StatsRequestResult{
		TotalEntries: int32(total.TotalEntries),
		First:        total.First,
		Latest:       total.Latest,
		JSONLog:      total.JSONLog,
	}, nil
}

func CreateStatsRequest(e *eventlog.Eventlog) (*StatsRequest, error) {
	return &StatsRequest{
		eventlog: e,
	}, nil
}
