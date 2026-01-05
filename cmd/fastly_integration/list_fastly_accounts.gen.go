package fastly_integration

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListFastlyAccountsCmd = &cobra.Command{
	Use: "list-fastly-accounts",

	Short: "List Fastly accounts",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		res, _, err := api.ListFastlyAccounts(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-fastly-accounts")

		cmd.Println(cmdutil.FormatJSON(res, "fastly-accounts"))
	},
}

func init() {
	Cmd.AddCommand(ListFastlyAccountsCmd)
}
