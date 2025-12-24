package fastly_integration

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListFastlyAccountsCmd = &cobra.Command{
	Use:   "listfastlyaccounts",
	Short: "List Fastly accounts",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		res, _, err := api.ListFastlyAccounts(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listfastlyaccounts: %v", err)
		}

		cmdutil.PrintJSON(res, "fastly_integration")
	},
}

func init() {
	Cmd.AddCommand(ListFastlyAccountsCmd)
}
