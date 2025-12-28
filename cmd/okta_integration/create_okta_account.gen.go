package okta_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateOktaAccountCmd = &cobra.Command{
	Use:   "create_okta_account",
	Short: "Add Okta account",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOktaIntegrationApi(client.NewAPIClient())
		res, _, err := api.CreateOktaAccount(client.NewContext(apiKey, appKey, site), datadogV2.OktaAccountRequest{})
		if err != nil {
			log.Fatalf("failed to create_okta_account: %v", err)
		}

		cmdutil.PrintJSON(res, "okta-accounts")
	},
}

func init() {
	Cmd.AddCommand(CreateOktaAccountCmd)
}
