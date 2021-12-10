package http

import (
	"github.com/debridge-finance/orbitdb-go/pkg/context"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
)

const (
	contentTypeHeader = "Content-Type"

	RequestInvalidContextKey RequestContextKey = iota
	RequestLoggerContextKey
	RequestIdContextKey
	PaginationContextKey
)

type RequestContextKey int

//

type Error struct {
	Error string `json:"error" swag_description:"error message" swag_example:"division by zero"`
}

func NewError(msg string) *Error {
	return &Error{Error: msg}
}

//

func Write(w ResponseWriter, r *Request, status int, payload interface{}) {
	w.Header().Set(contentTypeHeader, "application/json") // FIXME: other content types?
	w.WriteHeader(status)
	if payload != nil {
		err := Encode(r, w, payload)
		if err != nil {
			Log(r).
				Err(err).
				Msg("failed to encode result into client connection")
		}
	}
}

func WritePaginated(w ResponseWriter, r *Request, status int, payload interface{}) {
	pagination, ok := context.Get(r.Context(), PaginationContextKey).(Pagination)

	if !ok {
		Write(w, r, status, payload)
	}

	paginatedPayload := PaginationResult{
		Pagination: pagination,
		Result:     payload,
	}

	Write(w, r, status, paginatedPayload)

}

//

func WriteOk(w ResponseWriter, r *Request, payload interface{}) {
	Write(w, r, StatusOk, payload)
}

func WriteNotFound(w ResponseWriter, r *Request, payload interface{}) {
	Write(w, r, StatusNotFound, payload)
}

//

func WriteErrorMsg(w ResponseWriter, r *Request, code int, err error, msg string) {
	Log(r).
		Err(err).
		Int("code", code).
		Str("statusText", StatusText(code)).
		Msg("http handler error")
	if msg != "" {
		Write(w, r, code, NewError(msg))
	}
}

func WriteError(w ResponseWriter, r *Request, code int, err error) {
	var msg string
	if err != nil {
		msg = err.Error()
	}
	WriteErrorMsg(w, r, code, err, msg)
}

//

func Log(r *Request) *log.Logger {
	// XXX: because we use struct value and `l, _` this will not panic
	// it would always contain a logger, even if we have no logger middleware
	// in the chain(and no RequestLoggerContextKey value).
	// Never miss a message in the logs :)
	l, _ := r.Context().Value(RequestLoggerContextKey).(log.Logger)
	return &l
}
