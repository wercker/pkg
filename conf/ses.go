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

// SESOptions are the commonly used options when using the AWS SES service.
type SESOptions struct {
	SESRegion string
}

// ParseSESOptions fetches the values from urfave/cli Context and returns
// them as a SESOptions. Uses the names as specified in SESFlags.
func ParseSESOptions(c *cli.Context) *SESOptions {
	return &SESOptions{
		SESRegion: c.String("ses-region"),
	}
}

// SESFlags returns the flags that will be used by ParseSESOptions.
func SESFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "ses-region",
			Usage:  "AWS region",
			Value:  "us-east-1",
			EnvVar: "AWS_REGION",
		},
	}
}
