/*
 * RTS: Request to Struct. Generates Go structs from a server response.
 * Copyright (C) 2016-2022 Paolo Galeone <nessuno@nerdz.eu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 * Exhibit B is not attached; this software is compatible with the
 * licenses expressed under Section 1.12 of the MPL v2.
 */

package rts

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"unicode"

	"github.com/ChimeraCoder/gojson"
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

// replaceParameters replace parameters in the route line if present
// it returns the path with parameter replaced and the path with parameter
// (without ':') to build the name
func replaceParameters(line string) (string, string) {
	line = strings.TrimSpace(line)

	// if no parameters
	if !strings.Contains(line, " ") {
		return line, line
	}

	var ret bytes.Buffer

	pathAndParams := strings.Split(line, " ")
	parametersFound := 0
	path := pathAndParams[0]
	for i, l := 0, len(path); i < l; i++ {
		if path[i] == ':' {
			for ; i < l && path[i] != '/'; i++ {
			}
			// replace parameter with its value
			ret.WriteString(pathAndParams[1+parametersFound])
			parametersFound++
			if i < l && path[i] == '/' {
				ret.WriteRune('/')
			}
		} else {
			ret.WriteByte(path[i])
		}
	}

	return ret.String(), strings.Replace(path, ":", "", -1)
}

func requestConverter(server, line, pkg string, headerMap map[string]string, c chan result, wg *sync.WaitGroup, insecure, subStruct bool) {
	// Decrement the counter when goroutine ends
	defer wg.Done()

	requestPath, parametricRequest := replaceParameters(line)

	req, _ := http.NewRequest("GET", server+requestPath, nil)
	// set headers
	for key, value := range headerMap {
		req.Header.Set(key, value)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
	}

	client := &http.Client{Transport: tr}
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
			fmt.Errorf("Request: %s%s returned status %d\n", server, requestPath, res.StatusCode),
		}
		return
	}

	// Generate structName
	requestPathElements := strings.Split(parametricRequest, "/")
	var structName string
	for _, elem := range requestPathElements {
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
	tagList := []string{"json"}
	convertFloats := true
	r.res, r.err = gojson.Generate(res.Body, gojson.ParseJson, structName, pkg, tagList, subStruct, convertFloats)
	c <- r
}

func jsonToStruct(pkg, json string) result {

	mutex.Lock()
	unnamedStruct++
	structName := "Foo" + strconv.Itoa(unnamedStruct)
	mutex.Unlock()

	var r result
	tagList := []string{"json"}
	convertFloats := true
	subStruct := true
	r.res, r.err = gojson.Generate(strings.NewReader(json), gojson.ParseJson, structName, pkg, tagList, subStruct, convertFloats)
	return r
}
