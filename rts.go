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
	"go/format"
	"strings"
	"sync"
)

// Do exexutes the GET request to every route defined concatenated with server name.
// lines is a slice that contains routes in the format: `/path/to/request`
// or `/path/:parameter1/:parameter2/request parameter1Value parameter2Value`
// It passes headers in every request and returns a file whose package is pkg containing
// the struct definitions.
func Do(pkg, server string, lines []string, headerMap map[string]string) ([]byte, error) {
	server = strings.TrimRight(server, "/")
	lines = deleteEmpty(lines)

	var wg sync.WaitGroup
	n := len(lines)
	wg.Add(n)
	c := make(chan result, n)
	defer close(c)

	for i := 0; i < n; i++ {
		go requestConverter(server, lines[i], pkg, headerMap, c, &wg)
	}
	wg.Wait()

	var structs []byte
	for i := 0; i < n; i++ {
		r := <-c
		if r.err != nil {
			return r.res, r.err
		}
		structs = append(structs, r.res...)
	}

	fileContent := string(structs)
	fileContent = strings.Replace(fileContent, "}\npackage "+pkg, "}", -1)

	return format.Source([]byte(fileContent))
}
