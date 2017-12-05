//-----------------------------------------------------------------------------
// Copyright (c) 2017 Oracle and/or its affiliates.  All rights reserved.
// This program is free software: you can modify it and/or redistribute it
// under the terms of:
//
// (i)  the Universal Permissive License v 1.0 or at your option, any
//      later version (http://oss.oracle.com/licenses/upl); and/or
//
// (ii) the Apache License v 2.0. (http://www.apache.org/licenses/LICENSE-2.0)
//-----------------------------------------------------------------------------

package reflectutil

import (
	"reflect"
)

// GetMethods uses the reflect package to get the method names on defined on
// in.
func GetMethods(in interface{}) []string {
	if in == nil {
		return []string{}
	}

	t := reflect.TypeOf(in)
	if t.Kind() != reflect.Ptr {
		t = reflect.PtrTo(t)
	}

	numMethods := t.NumMethod()
	methods := make([]string, numMethods)
	for i := 0; i < numMethods; i++ {
		methods[i] = t.Method(i).Name
	}

	return methods
}
