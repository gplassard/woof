package fastly_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateFastlyAccountCmd = &cobra.Command{
	Use:   "create_fastly_account",
	Short: "Add Fastly account",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		res, _, err := api.CreateFastlyAccount(client.NewContext(apiKey, appKey, site), datadogV2.FastlyAccountCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create_fastly_account: %v", err)
		}

		cmdutil.PrintJSON(res, "fastly_integration")
	},
}

func init() {
	Cmd.AddCommand(CreateFastlyAccountCmd)
}
