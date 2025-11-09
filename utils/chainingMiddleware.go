package utils

import "net/http"

func ChainingMiddlewares(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	// aita te hbe nah krn aita flow tik thake nah
	// for _, m := range middlewares {
	// 	h = m(h)
	// }

	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}

	return h
}
