package main

import (
	"github.com/alexzhc/rmq/pkg/cli/broker"
	"github.com/alexzhc/rmq/pkg/cli/group"
	"github.com/alexzhc/rmq/pkg/cli/message"
	"github.com/alexzhc/rmq/pkg/cli/namesrv"
	"github.com/alexzhc/rmq/pkg/cli/topic"
	"github.com/alexzhc/rmq/pkg/config"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	root := &cobra.Command{
		Use:   "rmq",
		Short: "Apache RocketMQ cli",
	}

	r := new(config.RocketMQConfig)
	pf := root.PersistentFlags()
	pf.StringVar(&r.ConfigFile, "config", "", "config file")
	root.AddCommand(topic.NewCommand(r))
	root.AddCommand(group.NewCommand(r))
	root.AddCommand(message.NewCommand(r))
	root.AddCommand(namesrv.NewCommand(r))
	root.AddCommand(broker.NewCommand(r))

	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}

}
