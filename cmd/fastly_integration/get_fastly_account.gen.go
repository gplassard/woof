package fastly_integration

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetFastlyAccountCmd = &cobra.Command{
	Use:   "get-fastly-account [account_id]",
	
	Short: "Get Fastly account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		res, _, err := api.GetFastlyAccount(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-fastly-account")

		cmdutil.PrintJSON(res, "fastly-accounts")
	},
}

func init() {
	Cmd.AddCommand(GetFastlyAccountCmd)
}
