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

package log

import (
	"os"
	"time"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/sirupsen/logrus"
	cli "gopkg.in/urfave/cli.v1"
)

func SetupLogging(c *cli.Context) error {
	if c.GlobalBool("debug") {
		SetLevel(DebugLevel)
	}

	// Dynamically return false or true based on the logger output's
	// file descriptor referring to a terminal or not.
	if os.Getenv("TERM") == "dumb" || !isLogrusTerminal() {
		SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339Nano,
		})
	}
	return nil
}

// isLogrusTerminal checks if the standard logger of Logrus is a terminal.
func isLogrusTerminal() bool {
	w := logrus.StandardLogger().Out
	switch v := w.(type) {
	case *os.File:
		return terminal.IsTerminal(int(v.Fd()))
	default:
		return false
	}
}
