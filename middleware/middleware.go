package middleware

import "net/http"

type Middleware func(next http.Handler) http.Handler

// ChainMiddlewares handles all the middleware and chains the call for the actual api call
func ChainMiddlewares(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}
