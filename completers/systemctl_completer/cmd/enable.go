package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var enableCmd = &cobra.Command{
	Use:     "enable",
	Short:   "Enable one or more unit files",
	GroupID: "unit file",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(enableCmd).Standalone()

	rootCmd.AddCommand(enableCmd)

	carapace.Gen(enableCmd).PositionalAnyCompletion(
		carapace.Batch(
			action.ActionUnits(enableCmd).FilterArgs(),
			carapace.ActionFiles(),
		).ToA(),
	)
}
