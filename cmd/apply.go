package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/updatecli/updatecli/pkg/core/engine/manifest"
	"github.com/updatecli/updatecli/pkg/core/format"

	"github.com/spf13/cobra"
)

var (
	applyCommit bool
	applyClean  bool
	applyPush   bool

	applyCmd = &cobra.Command{
			/*
				Technically speaking we could have multiple policies to apply from the command line
				but for clarity we will only allow one policy at a time.
				This decision could be changed in the future.

				The same rule would also apply to the 'diff' and show subcommand.
			*/
			Args:  cobra.MatchAll(cobra.MaximumNArgs(1)),
			Use:   "apply NAME[:TAG|@DIGEST]",
			Short: "apply checks if an update is needed then apply the changes",
			Run: func(cmd *cobra.Command, args []string) {
				policyReferences = args
				err := getPolicyFilesFromRegistry()
				if err != nil {
					format.PrintError("❌ Command failed")
					logrus.Errorf("command failed: %s", err)
					os.Exit(1)
				}

				e.Options.Manifests = append(e.Options.Manifests, manifest.Manifest{
					Manifests: manifestFiles,
					Values:    valuesFiles,
					Secrets:   secretsFiles,
				})

				format.PrintTitle("🚀 Starting Apply")
				err = run("apply")
				if err != nil {
					format.PrintError("❌ Apply Failed")
					logrus.Errorf("command failed: %s", err)
					os.Exit(1)
				}
				format.PrintSuccess("✅ Apply Completed Successfully")
			},
		}
	)

func init() {
	applyCmd.Flags().StringArrayVarP(&manifestFiles, "config", "c", []string{}, "Sets config file or directory. By default, Updatecli looks for a file named 'updatecli.yaml' or a directory named 'updatecli.d'")
	applyCmd.Flags().StringVar(&udashOAuthAudience, "reportAPI", "", "Set the report API URL where to publish pipeline reports")
	applyCmd.Flags().StringArrayVarP(&valuesFiles, "values", "v", []string{}, "Sets values file uses for templating")
	applyCmd.Flags().StringArrayVar(&secretsFiles, "secrets", []string{}, "Sets Sops secrets file uses for templating")

	applyCmd.Flags().BoolVarP(&applyCommit, "commit", "", true, "Record changes to the repository, '--commit=false'")
	applyCmd.Flags().BoolVar(&applyClean, "clean", false, "Remove updatecli working directory like '--clean=true'")
}
