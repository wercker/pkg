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
