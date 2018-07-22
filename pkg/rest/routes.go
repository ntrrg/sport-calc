// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package rest

import (
	"net/http"

	nthttp "github.com/ntrrg/ntgo/net/http"
)

func Mux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/", HomeMux())
	// mux.Handle("/", http.FileServer(http.Dir("")))

	return mux
}

func HomeMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/", nthttp.AdaptFunc(Home, nthttp.JSONResponse()))

	return mux
}
