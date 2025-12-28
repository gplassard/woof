package fastly_integration

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListFastlyAccountsCmd = &cobra.Command{
	Use:   "list-fastly-accounts",
	
	Short: "List Fastly accounts",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		res, _, err := api.ListFastlyAccounts(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-fastly-accounts")

		cmdutil.PrintJSON(res, "fastly-accounts")
	},
}

func init() {
	Cmd.AddCommand(ListFastlyAccountsCmd)
}
