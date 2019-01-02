package httprouter

import (
	"net/http"
)

const (
	// paramsKey is the request context key under which URL params are stored.
	paramsKey = iota
)

// Params is used to save the URL parameters , as returned by the router.
type Params map[string]string

// ParamsFrom pulls the URL parameters from a request context
func ParamsFrom(r *http.Request) (p Params) {
	if r != nil {
		p, _ = r.Context().Value(paramsKey).(Params)
	}
	return
}
