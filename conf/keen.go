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
