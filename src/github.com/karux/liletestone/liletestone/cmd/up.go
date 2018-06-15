package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/karux/liletestone/subscribers"
	"github.com/lileio/lile"
	"github.com/lileio/pubsub"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "up runs both RPC service",
	Run: func(cmd *cobra.Command, args []string) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		go func() {

			//TODO: bootstrap secrets and service credentials  .
			lile.Run()
		}()
		go func() {
			pubsub.Subscribe(&subscribers.LiletestoneServiceSubscriber{})
		}()

		<-c
		lile.Shutdown()
		pubsub.Shutdown()
	},
}

func init() {
	RootCmd.AddCommand(upCmd)
}
