package log

import (
	cli "gopkg.in/urfave/cli.v1"
)

func SetupLogging(c *cli.Context) error {
	if c.GlobalBool("debug") {
		SetLevel(DebugLevel)
	}
	return nil
}
