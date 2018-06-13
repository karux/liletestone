package main

import (
	_ "net/http/pprof"

	"github.com/karux/liletestone"
	"github.com/karux/liletestone/liletestone/cmd"
	"github.com/karux/liletestone/server"
	"github.com/lileio/fromenv"
	"github.com/lileio/lile"
	"github.com/lileio/logr"
	"github.com/lileio/pubsub"
	"github.com/lileio/pubsub/middleware/defaults"
	"google.golang.org/grpc"
)

func main() {
	logr.SetLevelFromEnv()
	s := &server.LiletestoneServer{}

	lile.Name("liletestone")
	lile.Server(func(g *grpc.Server) {
		liletestone.RegisterLiletestoneServer(g, s)
	})

	pubsub.SetClient(&pubsub.Client{
		ServiceName: lile.GlobalService().Name,
		Provider:    fromenv.PubSubProvider(),
		Middleware:  defaults.Middleware,
	})

	cmd.Execute()
}
