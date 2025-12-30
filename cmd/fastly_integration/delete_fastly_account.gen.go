package fastly_integration

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteFastlyAccountCmd = &cobra.Command{
	Use: "delete-fastly-account [account_id]",

	Short: "Delete Fastly account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		_, err := api.DeleteFastlyAccount(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-fastly-account")

	},
}

func init() {
	Cmd.AddCommand(DeleteFastlyAccountCmd)
}
