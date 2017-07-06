package gateway

import (
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/wercker/auth/middleware"
	"github.com/wercker/pkg/config"
	"github.com/wercker/pkg/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	cli "gopkg.in/urfave/cli.v1"
)

var (
	// errorExitCode returns a urfave decorated error which indicates a exit
	// code 1. To be return from a urfave action.
	errorExitCode = cli.NewExitError("", 1)
)

type Registrar func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error

func Command(cfg *config.Config, action func(c *cli.Context) error) cli.Command {
	port := cfg.Int("gateway.port")
	grpcPort := cfg.Int("gprc.port")

	cmd := cli.Command{
		Name:   "gateway",
		Usage:  "Starts environment variable HTTP->gRPC gateway",
		Action: action,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:   "port, p",
				Value:  port,
				EnvVar: "HTTP_PORT",
			},
			cli.StringFlag{
				Name:   "host",
				Value:  fmt.Sprintf("localhost:%d", grpcPort),
				EnvVar: "GRPC_HOST",
			},
		},
	}
	return cmd
}

func Action(cfg *config.Config, reg func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error, muxOptions ...runtime.ServeMuxOption) func(c *cli.Context) error {
	var gatewayAction = func(c *cli.Context) error {
		log.Info("Starting gateway")

		log.Debug("Parsing gateway options")
		o, err := parseGatewayOptions(c)
		if err != nil {
			log.WithError(err).Error("Unable to validate arguments")
			return errorExitCode
		}

		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		mux := runtime.NewServeMux(muxOptions...)

		opts := []grpc.DialOption{grpc.WithInsecure()}
		err = reg(ctx, mux, o.Host, opts)
		if err != nil {
			log.WithError(err).Error("Unable to register handler from Endpoint")
			return errorExitCode
		}

		authMiddleware := middleware.AuthTokenMiddleware(mux)

		log.Infof("Listening on port %d", o.Port)
		http.ListenAndServe(fmt.Sprintf(":%d", o.Port), authMiddleware)

		return nil
	}
	return gatewayAction
}

func parseGatewayOptions(c *cli.Context) (*gatewayOptions, error) {
	port := c.Int("port")
	// if !validPortNumber(port) {
	//   return nil, fmt.Errorf("Invalid port number: %d", port)
	// }

	return &gatewayOptions{
		Port: port,
		Host: c.String("host"),
	}, nil
}

type gatewayOptions struct {
	Port int
	Host string
}
