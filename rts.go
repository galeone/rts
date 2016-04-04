/*
RTS: Request to Struct. Generates Go structs from a server response.
Copyright (C) 2016 Paolo Galeone <nessuno@nerdz.eu>

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
Exhibit B is not attached; this software is compatible with the
licenses expressed under Section 1.12 of the MPL v2.
*/

// Package rts generates Go structs from JSON server responses.
// It defines type names using the specified routes and skipping numbers
// (e.g: a request to a route like /users/1/posts generates type UsersPosts)
// Supports headers pesonalization as well, thus RTS can be used to generate
// types from response protected by some authorization method
package rts

import (
	"fmt"
	"github.com/ChimeraCoder/gojson"
	"github.com/nerdzeu/nerdz-api/utils"
	"go/format"
	"net/http"
	"strconv"
	"strings"
)

// Do exexutes the GET request to every route defined concatenated with server name.
// It passes headers in every request and returns a file whose package is pkg containing
// the struct definitions.
func Do(pkg, server string, routes []string, headerMap map[string]string) ([]byte, error) {
	var structs []byte
	client := &http.Client{}
	unnamedStruct := 0
	var res *http.Response
	var e error
	for _, route := range routes {
		if route != "" {
			req, _ := http.NewRequest("GET", server+route, nil)
			// set headers
			for key, value := range headerMap {
				req.Header.Set(key, value)
			}

			if res, e = client.Do(req); e != nil {
				return nil, e
			}
			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				return nil, fmt.Errorf("Request: %s%s returned status %d\n", server, route, res.StatusCode)
			}

			// Generate structName
			routeElements := strings.Split(route, "/")
			var structName string
			for _, elem := range routeElements {
				_, e = strconv.ParseInt(elem, 10, 64)
				// Skip numbers
				if e == nil {
					continue
				}
				structName += utils.UpperFirst(elem)
			}
			if structName == "" {
				unnamedStruct++
				structName = "Foo" + strconv.Itoa(unnamedStruct)
			}
			var buff []byte
			if buff, e = json2struct.Generate(res.Body, structName, pkg); e != nil {
				return nil, e
			}
			structs = append(structs, buff...)
		}
	}

	fileContent := string(structs)
	fileContent = strings.Replace(fileContent, "}\npackage "+pkg, "}", -1)

	return format.Source([]byte(fileContent))
}
