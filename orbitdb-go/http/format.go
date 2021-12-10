package http

import (
	"encoding/json"

	"github.com/debridge-finance/orbitdb-go/pkg/io"
)

// this Encode and Decode are just abstractions which incapsulate json encoding
// other formats could be implemented in future, this is why you should provide a Request

func Encode(req *Request, w io.Writer, v interface{}) error { return json.NewEncoder(w).Encode(v) }

func Decode(req *Request, v interface{}, r io.Reader) error { return json.NewDecoder(r).Decode(v) }
