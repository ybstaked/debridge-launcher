package main

import (
	"github.com/debridge-finance/orbitdb-go/http"
	"github.com/debridge-finance/orbitdb-go/http/spec"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
)

//

func hello(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	http.Write(w, r, http.StatusOk, "hello")
}

//

var dogs = Dogs{}

func postDogs(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	dog := Dog{}

	err := http.Decode(r, &dog, r.Body)
	if err != nil {
		panic(err)
	}

	dogs = append(dogs, &dog)

	http.Write(w, r, http.StatusCreated, nil)
}

func getDogs(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	http.Write(w, r, http.StatusOk, dogs)
}

type Dog struct {
	Name    string `json:"name" swag_example:"rex" swag_description:"dog name"`
	Madness int    `json:"madness" swag_example:"50" swag_description:"dog madness level"`
}

type Dogs = []*Dog

//

func getError(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	http.WriteErrorMsg(
		w, r, http.StatusInternalServerError,
		errors.New("error describing what goes wrong for developers/sysops"),
		"error describing what goes wrong for a user",
	)
}

//

func main() {
	l, err := log.Create(log.DefaultConfig)
	if err != nil {
		panic(err)
	}

	ms, err := http.CreateMiddlewares(
		http.DefaultMiddlewareRegistry,
		http.DefaultMiddlewareConfig,
		l,
	)
	if err != nil {
		panic(err)
	}

	es := spec.Endpoints{
		spec.NewEndpoint("get", "/hello", "Say hello",
			spec.EndpointHandler(hello),
			spec.EndpointDescription("This endpoint returns a single string `hello`"),
			spec.EndpointResponse(http.StatusOk, "", "Successfully retrieved data"),
			spec.EndpointProduces("application/json"),
		),
		spec.NewEndpoint("post", "/dogs", "Add new dog to the registry",
			spec.EndpointHandler(postDogs),
			spec.EndpointDescription("This endpoint receives new Dog entity and saves it inside inmemory collection"),
			spec.EndpointBody(Dog{}, "Dog record", true),
			spec.EndpointResponse(http.StatusCreated, "", "Successfully retrieved and saved data"),
			spec.EndpointConsumes("application/json"),
		),
		spec.NewEndpoint("get", "/dogs", "Get dogs registry",
			spec.EndpointHandler(getDogs),
			spec.EndpointDescription("This endpoint returns whole registry with dogs records"),
			spec.EndpointResponse(http.StatusOk, Dogs{}, "Successfully retrieved data"),
			spec.EndpointProduces("application/json"),
		),
		spec.NewEndpoint("get", "/error", "Fails with some error",
			spec.EndpointHandler(getError),
			spec.EndpointDescription("This endpoint fails with internal server error leaving a log entry and sends a user custom message"),
			spec.EndpointResponse(http.StatusInternalServerError, "", "An error which is always returned"),
			spec.EndpointProduces("application/json"),
		),
	}

	err = http.New(
		http.DefaultConfig, l,
		ms, es,
		spec.Title("example server"),
	).ListenAndServe()
	if err != nil {
		panic(err)
	}
}
