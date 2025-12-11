package topic

import (
	"github.com/alexzhc/rmq/pkg/config"
	"github.com/spf13/cobra"
)

func NewCommand(r *config.RocketMQConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "topic",
		Short: "Create, delete, produce to and consume RocketMQ topics",
	}
	r.InstallRocketMQFlags(cmd)

	cmd.AddCommand(
		Create(r),
		Produce(r),
		List(r),
		Describe(r),
	)
	return cmd
}
