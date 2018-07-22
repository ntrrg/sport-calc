// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package rest

import (
	"net/http"

	nthttp "github.com/ntrrg/ntgo/net/http"
)

func Mux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle(
		"/forms/",
		http.StripPrefix("/forms", http.FileServer(http.Dir("static"))),
	)

	mux.Handle("/teams/", http.StripPrefix("/teams", TeamsMux()))

	return mux
}

func TeamsMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/", nthttp.AdaptFunc(Teams, nthttp.JSONResponse()))

	return mux
}
