// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package rest

import (
	"net/http"

	nthttp "github.com/ntrrg/ntgo/net/http"
	"github.com/ntrrg/ntgo/net/http/middleware"
)

func Mux(server *nthttp.Server) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("static")))

	mux.Handle("/teams/", middleware.Adapt(
		TeamsMux(),
		middleware.StripPrefix("/teams"),
		middleware.JSONResponse(),
		middleware.Cache("max-age=3600, s-max-age=3600"),
		middleware.Gzip(-1),
	))

	return mux
}

func TeamsMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Teams)

	return mux
}
