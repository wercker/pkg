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

import (
	"encoding/hex"

	cli "gopkg.in/urfave/cli.v1"
)

// AuthClientOptions are the commonly used options when using a AuthClient.
type AuthClientOptions struct {
	AuthTarget string
	ServiceKey []byte
}

// ParseAuthClientOptions fetches the values from urfave/cli Context and
// returns them as a AuthClientOptions. Uses the names as specified in
// AuthClientFlags.
func ParseAuthClientOptions(c *cli.Context) *AuthClientOptions {
	decodedServiceKey := []byte{}
	serviceKey := c.String("service-key")
	if serviceKey != "" {
		decodedServiceKey, _ = hex.DecodeString(serviceKey)
	}

	return &AuthClientOptions{
		AuthTarget: c.String("auth"),
		ServiceKey: decodedServiceKey,
	}
}

// AuthClientFlags returns the flags that will be used by
// ParseAuthClientOptions.
func AuthClientFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "auth",
			Value:  "localhost:6002",
			Usage:  "host and port of auth service",
			EnvVar: "AUTH_TARGET",
		},
		cli.StringFlag{
			Name:   "service-key",
			Usage:  "Hex encoded service key to use",
			EnvVar: "WERCKER_SERVICE_KEY",
		},
	}
}
