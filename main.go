// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package main

import (
	"flag"
	"log"

	nthttp "github.com/ntrrg/ntgo/net/http"

	"github.com/ntrrg/sport-calc/pkg/rest"
)

var cfg nthttp.Config

func main() {
	srv := nthttp.NewServer(cfg)

	srv.Handler = rest.Mux(srv)
	log.Println(srv.ListenAndServe())
	<-srv.Done
}

func init() {
	flag.StringVar(
		&cfg.Addr,
		"addr",
		":4000",
		"TCP address to listen on. If a path to a file is given, the server will" +
		" use a Unix Domain Socket.",
	)

	flag.Parse()
}
