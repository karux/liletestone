package main

import (
	"fmt"
	_ "net/http/pprof"

	"github.com/karux/go-utils/security"
	"github.com/karux/liletestone"
	"github.com/karux/liletestone/liletestone/cmd"
	"github.com/karux/liletestone/server"
	"github.com/lileio/fromenv"
	"github.com/lileio/lile"
	"github.com/lileio/logr"
	"github.com/lileio/pubsub"
	"github.com/lileio/pubsub/middleware/defaults"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	logr.SetLevelFromEnv()
	s := &server.LiletestoneServer{}

	lile.Name("liletestone")
	lile.Server(func(g *grpc.Server) {
		creds, err := credentials.NewServerTLSFromFile("./keystore/server/server-cert.pem", "./keystore/server/server-key.pem")
		if err == nil {
			fmt.Println("saving creds", creds)
			// lile.GlobalService().GRPCOptions
			sOptions := lile.GlobalService().GRPCOptions
			lile.GlobalService().GRPCOptions = append(sOptions, grpc.Creds(creds))
			fmt.Println("options", lile.GlobalService().GRPCOptions)
		}

		// Add security interceptor
		lile.AddUnaryInterceptor(security.NewAuthTokenServerInterceptor())

		//TODO: Add interceptor to customize Context
		// custom context = zap GetLogger,

		liletestone.RegisterLiletestoneServer(g, s)
	})

	pubsub.SetClient(&pubsub.Client{
		ServiceName: lile.GlobalService().Name,
		Provider:    fromenv.PubSubProvider(),
		Middleware:  defaults.Middleware,
	})

	cmd.Execute()
}
