package spec

import (
	"net/http"

	"github.com/rakyll/statik/fs"
)

const (
	SpecUIPath = "/swagger"
	SpecPath   = "/swagger.json"
)

func NewSwaggerUIHandler(prefix string) http.Handler {
	f, err := fs.New()
	if err != nil {
		panic(err)
	}

	return http.StripPrefix(
		prefix+SpecUIPath+"/",
		http.FileServer(f),
	)
}

func NewSwaggerEndpoints(specHandler http.Handler, uiHandler http.Handler) []*Endpoint {
	return []*Endpoint{
		NewEndpoint("get", SpecPath, "View this specification",
			EndpointHandler(specHandler),
			EndpointDescription("This endpoint serves a service API specification"),
			EndpointConsumes("application/json"),
		),
		NewEndpoint("get", SpecUIPath, "SwaggerUI redirector",
			EndpointHandler(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Add("Location", r.URL.Path+"/")
				w.WriteHeader(http.StatusFound)
			}),
			EndpointDescription("This endpoint redirects a user to the valid directory index to serve SwaggerUI"),
		),
		NewEndpoint("get", SpecUIPath+"/", "SwaggerUI",
			EndpointHandler(uiHandler),
			EndpointDescription("Serves SwaggerUI"),
		),
	}
}
