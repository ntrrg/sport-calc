// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package rest

import (
	"encoding/json"
	"net/http"

	"github.com/ntrrg/sport-calc/pkg/api"
)

func Teams(w http.ResponseWriter, req *http.Request) {
	teams := []api.Team {
		{Name: "Millan"},
		{Name: "EMBCE"},
	}

	json.NewEncoder(w).Encode(teams)
}
