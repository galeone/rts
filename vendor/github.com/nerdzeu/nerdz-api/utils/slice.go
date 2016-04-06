/*
RtS: Request to Struct. Generates Go structs from a server response.
Copyright (C) 2016 Paolo Galeone <nessuno@nerdz.eu>

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
Exhibit B is not attached; this software is compatible with the
licenses expressed under Section 1.12 of the MPL v2.
*/

/*
Copyright (C) 2016 Paolo Galeone <nessuno@nerdz.eu>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package utils

import (
	"fmt"
	"reflect"
)

// copyElement returns the reflect.Value, (deep) copy of source
func copyElement(source reflect.Value) reflect.Value {
	Type := source.Type()
	elem := reflect.Indirect(reflect.New(Type))

	switch Type.Kind() {
	case reflect.Struct:
		for j := 0; j < Type.NumField(); j++ {
			elem.Field(j).Set(source.Field(j))
		}
	default:
		elem.Set(source)
	}

	return elem
}

// ReverseSlice reverse slice if slice is of type reflect.Slice or reflrect.Ptr (to silce)
// panics if slice type is different
func ReverseSlice(slice interface{}) interface{} {
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice, reflect.Ptr:
		values := reflect.Indirect(reflect.ValueOf(slice))
		if values.Len() == 0 {
			return slice
		}

		reversedSlice := reflect.MakeSlice(reflect.SliceOf(reflect.Indirect(values.Index(0)).Type()), values.Len(), values.Len())

		k := 0
		for i := values.Len() - 1; i >= 0; i-- {
			reversedSlice.Index(k).Set(copyElement(values.Index(i)))
			k++
		}

		return reversedSlice.Interface()
	default:
		panic(fmt.Sprintf("Type not allowed: %v", reflect.TypeOf(slice).Kind()))
	}
}

// InSlice returns true if value is in slice
func InSlice(value, slice interface{}) bool {
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice, reflect.Ptr:
		values := reflect.Indirect(reflect.ValueOf(slice))
		if values.Len() == 0 {
			return false
		}

		val := reflect.Indirect(reflect.ValueOf(value))

		if val.Kind() != values.Index(0).Kind() {
			return false
		}

		for i := 0; i < values.Len(); i++ {
			if reflect.DeepEqual(values.Index(i).Interface(), val.Interface()) {
				return true
			}
		}
	}
	return false
}
