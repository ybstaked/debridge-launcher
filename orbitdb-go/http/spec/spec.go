package spec

import (
	"github.com/debridge-finance/orbitdb-go/pkg/reflect"
	"github.com/savaki/swag"
	"github.com/savaki/swag/endpoint"
	"github.com/savaki/swag/swagger"
)

type (
	Spec                    = swagger.API
	Option                  = swag.Option
	Options                 = []Option
	EndpointOption          = endpoint.Option
	EndpointBuilder         = endpoint.Builder
	EndpointResponseOption  = endpoint.ResponseOption
	EndpointResponseOptions = []EndpointResponseOption

	Endpoint  = swagger.Endpoint
	Endpoints = []*Endpoint
)

func New(options ...Option) *Spec { return swag.New(options...) }
func NewEndpoint(method, path, summary string, options ...EndpointOption) *Endpoint {
	return endpoint.New(method, path, summary, options...)
}

func EndpointsToOption(endpoint ...*Endpoint) Option { return swag.Endpoints(endpoint...) }

//

func BasePath(v string) Option     { return swag.BasePath(v) }
func ContactEmail(v string) Option { return swag.ContactEmail(v) }
func Description(v string) Option  { return swag.Description(v) }
func Host(v string) Option         { return swag.Host(v) }
func Schemes(v ...string) Option   { return swag.Schemes(v...) }
func Title(v string) Option        { return swag.Title(v) }
func Version(v string) Option      { return swag.Version(v) }

//

func EndpointBody(prototype interface{}, description string, required bool) EndpointOption {
	return endpoint.Body(prototype, description, required)
}
func EndpointBodyType(t reflect.Type, description string, required bool) EndpointOption {
	return endpoint.BodyType(t, description, required)
}
func EndpointConsumes(v ...string) EndpointOption        { return endpoint.Consumes(v...) }
func EndpointDescription(v string) EndpointOption        { return endpoint.Description(v) }
func EndpointHandler(handler interface{}) EndpointOption { return endpoint.Handler(handler) }
func EndpointOperationID(v string) EndpointOption        { return endpoint.OperationID(v) }
func EndpointPath(name, typ, description string, required bool) EndpointOption {
	return endpoint.Path(name, typ, description, required)
}
func EndpointProduces(v ...string) EndpointOption { return endpoint.Produces(v...) }
func EndpointQuery(name, typ, description string, required bool) EndpointOption {
	return endpoint.Query(name, typ, description, required)
}
func EndpointResponse(code int, prototype interface{}, description string, opts ...EndpointResponseOption) EndpointOption {
	return endpoint.Response(code, prototype, description, opts...)
}
func EndpointResponseType(code int, t reflect.Type, description string, opts ...EndpointResponseOption) EndpointOption {
	return endpoint.ResponseType(code, t, description, opts...)
}
func EndpointTags(tags ...string) EndpointOption { return endpoint.Tags(tags...) }

func EndpointPaginationParameters() EndpointOption {
	o := EndpointQuery("offset", "int", "Allows you to specify the ranking number of the first item on the page", false)
	l := EndpointQuery("limit", "int", "Allows you to set the number of objects returned in one page", false)
	return func(b *EndpointBuilder) {
		o(b)
		l(b)
	}
}
