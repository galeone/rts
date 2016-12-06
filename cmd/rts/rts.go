/*
RTS: Request to Struct. Generates Go structs from a server response.
Copyright (C) 2016 Paolo Galeone <nessuno@nerdz.eu>

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
Exhibit B is not attached; this software is compatible with the
licenses expressed under Section 1.12 of the MPL v2.
*/

package main

import (
	"flag"
	"fmt"
	"github.com/galeone/rts"
	"io/ioutil"
	"os"
	"strings"
)

var (
	help       bool
	server     string
	pkgName    string
	routesFile string
	headers    string
	out        string
	insecure   bool
	subStruct  bool
)

func init() {
	flag.BoolVar(&help, "help", false, "prints this help")
	flag.StringVar(&server, "server", "http://localhost:9090", "sets the server address")
	flag.StringVar(&routesFile, "routes", "routes.txt", "Routes to request. One per line")
	flag.StringVar(&headers, "headers", "", "Headers to add in every request")
	flag.StringVar(&out, "out", "", "Output file. Stdout is used if not specified")
	flag.StringVar(&pkgName, "pkg", "main", "Package name")
	flag.BoolVar(&insecure, "insecure", false, "Disables TLS Certificate check for HTTPS, use in case HTTPS Server Certificate is signed by an unknown authority")
	flag.BoolVar(&subStruct, "substruct", false, "Creates types for sub-structs")
}

func main() {
	flag.Parse()
	if help {
		fmt.Println("usage: rts [options]")
		flag.PrintDefaults()
		return
	}

	// If the route file does not exists, it does a single request to
	// the server address
	content, _ := ioutil.ReadFile(routesFile)
	routes := strings.Split(string(content), "\n")

	if len(routes) == 1 && routes[0] == "" {
		routes = append(routes, "/")
	}

	// parse Headers
	headerArr := strings.Split(headers, "\n")
	var headerMap = make(map[string]string)
	for _, headerLine := range headerArr {
		if headerLine != "" {
			h := strings.SplitAfterN(headerLine, ":", 2)
			if len(h) != 2 {
				fmt.Printf("Malformed header. Expected <key>:<value> Got: %s\n", h[0])
				os.Exit(1)
			}
			headerMap[strings.Split(h[0], ":")[0]] = h[1]
		}
	}

	structs, e := rts.Do(pkgName, server, routes, headerMap, insecure, subStruct)
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}
	if out == "" {
		fmt.Print(string(structs))
		return
	}

	ioutil.WriteFile(out, structs, 0644)
}
