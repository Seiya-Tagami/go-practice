package utils

import "net/http"

type Middleware func(http.Handler) http.Handler

func ChainMiddlewares(h http.Handler, ms ...Middleware) http.Handler {
	if len(ms) < 1 {
		return h
	}

	wrappedHandler := h
	for i := len(ms) - 1; i >= 0; i-- {
		wrappedHandler = ms[i](wrappedHandler)
	}

	return wrappedHandler
}
