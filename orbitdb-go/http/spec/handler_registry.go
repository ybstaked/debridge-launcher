package spec

import (
	"net/http"
)

type HandlerRegistry map[string]http.Handler

func (r HandlerRegistry) Add(name string, h http.Handler) HandlerRegistry {
	r[name] = h
	return r
}

func (r HandlerRegistry) Del(name string) HandlerRegistry {
	delete(r, name)
	return r
}

func (r HandlerRegistry) Get(name string) http.Handler {
	return r[name]
}
