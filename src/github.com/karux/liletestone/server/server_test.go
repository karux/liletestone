package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/karux/liletestone"
	"github.com/lileio/lile"
)

var s = LiletestoneServer{}
var cli liletestone.LiletestoneClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		liletestone.RegisterLiletestoneServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = liletestone.NewLiletestoneClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
