// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package pkg

import (
	"net/http"

	nthttp "github.com/ntrrg/ntgo/net/http"
)

func Mux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/", HomeMux())

	return mux
}

func HomeMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/", nthttp.AdaptFunc(HomeHandler, nthttp.JSONResponse()))

	return mux
}
