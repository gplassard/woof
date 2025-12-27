package okta_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetOktaAccountCmd = &cobra.Command{
	Use:   "getoktaaccount [account_id]",
	Short: "Get Okta account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOktaIntegrationApi(client.NewAPIClient())
		res, _, err := api.GetOktaAccount(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getoktaaccount: %v", err)
		}

		cmdutil.PrintJSON(res, "okta_integration")
	},
}

func init() {
	Cmd.AddCommand(GetOktaAccountCmd)
}
