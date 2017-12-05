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

package conf

import "gopkg.in/urfave/cli.v1"

// ConcatFlags concatenates all flags provided to a single array. Currently it
// does not inspect if an flag already exists, or sort by name etc.
func ConcatFlags(flags ...[]cli.Flag) []cli.Flag {
	result := []cli.Flag{}

	for _, f := range flags {
		result = append(result, f...)
	}

	return result
}
