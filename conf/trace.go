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

// TraceOptions are the commonly used options when using a Trace service.
type TraceOptions struct {
	Trace         bool
	TraceEndpoint string
}

// ParseTraceOptions fetches the values from urfave/cli Context and
// returns them as a TraceOptions. Uses the names as specified in
// TraceFlags.
func ParseTraceOptions(c *cli.Context) *TraceOptions {
	return &TraceOptions{
		Trace:         c.Bool("trace"),
		TraceEndpoint: c.String("trace-endpoint"),
	}
}

// TraceFlags returns the flags that will be used by
// ParseTraceOptions.
func TraceFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:   "trace",
			Usage:  "Enable tracing",
			EnvVar: "TRACE_ENABLED",
		},
		cli.StringFlag{
			Name:   "trace-endpoint",
			Usage:  "Endpoint for the trace service",
			EnvVar: "TRACE_ENDPOINT",
		},
	}
}
