package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/updatecli/updatecli/pkg/core/format"
)

var (
	udashOAuthAccessToken string
	udashOAuthClientID    string
	udashOAuthIssuer      string
	udashOAuthAudience    string
	udashEndpointURL      string
	udashEndpointAPIURL   string

	udashLoginCmd = &cobra.Command{
		Use:     "login url",
		Short:   "[Experimental] login authenticates with the Udash.",
		Example: "updatecli udash login app.updatecli.io",
		Run: func(cmd *cobra.Command, args []string) {

			// TODO: To be removed once not experimental anymore
			if !experimental {
				format.PrintError("⚠️ The 'login' feature requires the flag experimental to work, such as:\n\t`updatecli udash login --experimental https://app.updatecli.io`")
				os.Exit(1)
			}

			switch len(args) {
			case 0:
				format.PrintError("❌ Missing URL to login to")
				os.Exit(1)
			case 1:
				udashEndpointURL = args[0]
			default:
				format.PrintError("❌ Can only login to one URL at a time")
				os.Exit(1)
			}

			format.PrintTitle("🔐 Starting Udash Login")
			err := run("udash/login")
			if err != nil {
				format.PrintError("❌ Command failed")
				os.Exit(1)
			}
			format.PrintSuccess("✅ Udash Login Completed Successfully")
		},
	}
)

func init() {
	udashLoginCmd.Flags().StringVar(&udashOAuthClientID, "oauth-clientId", "", "oauth-clientId defines the Oauth client ID")
	udashLoginCmd.Flags().StringVar(&udashOAuthIssuer, "oauth-issuer", "", "oauth-issuer defines the Oauth authentication URL")
	udashLoginCmd.Flags().StringVar(&udashOAuthAudience, "oauth-audience", "", "oauth-audience defines the Oauth audience URL")
	udashLoginCmd.Flags().StringVar(&udashOAuthAccessToken, "oauth-access-token", "", "oauth-access-token defines the Oauth access token")
	udashLoginCmd.Flags().StringVar(&udashEndpointAPIURL, "api-url", "", "api-url defines the udash API URL")

	udashCmd.AddCommand(udashLoginCmd)
}
