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
