/*
RTS: Request to Struct. Generates Go structs from a server response.
Copyright (C) 2016 Paolo Galeone <nessuno@nerdz.eu>

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
Exhibit B is not attached; this software is compatible with the
licenses expressed under Section 1.12 of the MPL v2.
*/

package rts

import (
	"fmt"
	"github.com/ChimeraCoder/gojson"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

var (
	mutex         sync.Mutex
	unnamedStruct int
)

type result struct {
	res []byte
	err error
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func requestConverter(server, route, pkg string, headerMap map[string]string, c chan result, wg *sync.WaitGroup) {
	// Decrement the counter when goroutine ends
	defer wg.Done()

	req, _ := http.NewRequest("GET", server+route, nil)
	// set headers
	for key, value := range headerMap {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	var e error
	var res *http.Response
	if res, e = client.Do(req); e != nil {
		c <- result{nil, e}
		return
	}
	// close writer on end
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		c <- result{
			nil,
			fmt.Errorf("Request: %s%s returned status %d\n", server, route, res.StatusCode),
		}
		return
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

		var firstFound bool
		for _, r := range elem {
			if unicode.IsLetter(r) {
				if !firstFound {
					structName += string(unicode.ToUpper(r))
					firstFound = true
				} else {
					structName += string(r)
				}
			}
		}
	}
	if structName == "" {
		mutex.Lock()
		unnamedStruct++
		structName = "Foo" + strconv.Itoa(unnamedStruct)
		mutex.Unlock()
	}

	var r result
	r.res, r.err = json2struct.Generate(res.Body, structName, pkg)
	c <- r
}
