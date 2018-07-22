// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package rest

import (
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	name := "Hello World"
	json.NewEncoder(w).Encode(Person{Name: name})
}
