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

// MongoOptions are the commonly used options when connecting to a MongoDB
// server.
type MongoOptions struct {
	MongoURI      string
	MongoDatabase string
}

// ParseMongoOptions fetches the values from urfave/cli Context and returns
// them as a MongoOptions. Uses the names as specified in MongoFlags.
func ParseMongoOptions(c *cli.Context) *MongoOptions {
	return &MongoOptions{
		MongoURI:      c.String("mongo"),
		MongoDatabase: c.String("mongo-database"),
	}
}

// MongoFlags returns the flags that will be used by ParseMongoOptions.
// defaultDatabase will be used for the --mongo-database flag.
func MongoFlags(defaultDatabase string) []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "mongo",
			Usage:  "MongoDB connection string",
			Value:  "mongodb://localhost:27017",
			EnvVar: "MONGODB_URI",
		},
		cli.StringFlag{
			Name:   "mongo-database",
			Usage:  "MongoDB Database",
			Value:  defaultDatabase,
			EnvVar: "MONGODB_DATABASE",
		},
	}
}
