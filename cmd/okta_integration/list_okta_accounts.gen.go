package okta_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListOktaAccountsCmd = &cobra.Command{
	Use:   "list-okta-accounts",
	Short: "List Okta accounts",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOktaIntegrationApi(client.NewAPIClient())
		res, _, err := api.ListOktaAccounts(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-okta-accounts: %v", err)
		}

		cmdutil.PrintJSON(res, "okta-accounts")
	},
}

func init() {
	Cmd.AddCommand(ListOktaAccountsCmd)
}
