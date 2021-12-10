package context

import (
	"context"
)

type (
	Context = context.Context
)

var (
	Background   = context.Background
	WithTimeout  = context.WithTimeout
	WithDeadline = context.WithDeadline
	WithValue    = context.WithValue
	Canceled     = context.Canceled
)

func Get(ctx Context, key interface{}) interface{} { return ctx.Value(key) }
func Put(ctx Context, key interface{}, value interface{}) Context {
	return context.WithValue(ctx, key, value)
}
