package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/updatecli/updatecli/pkg/core/format"

	"github.com/spf13/cobra"
)

var (
	manifestInitPolicyRootDir string
	manifestInitCmd           = &cobra.Command{
		Args:  cobra.MatchAll(cobra.MaximumNArgs(1)),
		Use:   "init <path>",
		Short: "init a new Updatecli policy",
		Run: func(cmd *cobra.Command, args []string) {

			manifestInitPolicyRootDir = "."
			if len(args) == 1 {
				manifestInitPolicyRootDir = args[0]
			}

			format.PrintTitle("🚀 Initializing Updatecli Policy")
			err := run("manifest/init")
			if err != nil {
				format.PrintError("❌ Command failed")
				logrus.Errorf("command failed: %s", err)
				os.Exit(1)
			}
			format.PrintSuccess("✅ Updatecli Policy Initialized Successfully")
		},
	}
)

func init() {
	manifestCmd.AddCommand(manifestInitCmd)
}
