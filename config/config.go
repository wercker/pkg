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

package config

import "strconv"

type Config struct {
	Data map[string]string
}

func (c *Config) String(k string) string {
	v, _ := c.Data[k]
	return v
}

func (c *Config) Int(k string) int {
	s, _ := c.Data[k]
	v, _ := strconv.Atoi(s)
	return v
}
