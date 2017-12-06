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

import cli "gopkg.in/urfave/cli.v1"

// KeenOptions are the commonly used options when sending metrics to Keen.
type KeenOptions struct {
	KeenProjectID string
	KeenWriteKey  string
}

// ParseKeenOptions fetches the values from urfave/cli Context and returns
// them as a KeenOptions. Uses the names as specified in KeenFlags.
func ParseKeenOptions(c *cli.Context) *KeenOptions {
	return &KeenOptions{
		KeenProjectID: c.String("keen-project-id"),
		KeenWriteKey:  c.String("keen-write-key"),
	}
}

// KeenFlags returns the flags that will be used by ParseKeenOptions.
func KeenFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "keen-project-id",
			Usage:  "Keen project ID to use when sending metrics to Keen",
			EnvVar: "KEEN_PROJECT_ID",
		},
		cli.StringFlag{
			Name:   "keen-write-key",
			Usage:  "Keen write key to use when sending metrics to Keen",
			EnvVar: "KEEN_WRITE_KEY",
		},
	}
}
