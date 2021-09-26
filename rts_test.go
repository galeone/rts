/*
 * RTS: Request to Struct. Generates Go structs from a server response.
 * Copyright (C) 2016-2021 Paolo Galeone <nessuno@nerdz.eu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 * Exhibit B is not attached; this software is compatible with the
 * licenses expressed under Section 1.12 of the MPL v2.
 */

package rts_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/galeone/rts"
)

var (
	server = "https://api.github.com"
	lines  = []string{
		"/",
		"/repos/:user/:repo galeone igor",
	}
	pkg                = "example"
	headerMap          = map[string]string{}
	expectedGeneration []byte
	insecure           = false
	subStruct          = false
)

func init() {
	var e error
	expectedGeneration, e = ioutil.ReadFile("expected_out")
	if e != nil {
		panic(e.Error())
	}
}

func TestDo(t *testing.T) {
	var file []byte
	var e error
	file, e = rts.Do(pkg, server, lines, headerMap, insecure, subStruct)
	if e != nil {
		t.Fatalf("No error expected but got: %s", e.Error())
	}

	if !bytes.Equal(file, expectedGeneration) {
		t.Errorf("Results should be equal, but they differs")
		t.Error(string(file))
	}
}
