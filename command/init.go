package command

import "github.com/spf13/cobra"

func NewCfCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cf [OPTIONS] COMMAND [arg...]",
		Short: "A command use for caiflower",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				cmd.Help()
				return
			}
		},
	}

	cmd.AddCommand(NewCreateCommand())
	return cmd
}
