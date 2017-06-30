package log

import (
	"os"

	"github.com/Sirupsen/logrus"
	cli "gopkg.in/urfave/cli.v1"
)

func SetupLogging(c *cli.Context) error {
	if c.GlobalBool("debug") {
		SetLevel(DebugLevel)
	}

	// Dynamically return false or true based on the logger output's
	// file descriptor referring to a terminal or not.
	if os.Getenv("TERM") == "dumb" || !logrus.IsTerminal(logrus.StandardLogger().Out) {
		SetFormatter(&logrus.JSONFormatter{})
	}
	return nil
}
