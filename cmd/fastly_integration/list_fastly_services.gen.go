package fastly_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListFastlyServicesCmd = &cobra.Command{
	Use:   "list-fastly-services [account_id]",
	Short: "List Fastly services",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		res, _, err := api.ListFastlyServices(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to list-fastly-services: %v", err)
		}

		cmdutil.PrintJSON(res, "fastly-services")
	},
}

func init() {
	Cmd.AddCommand(ListFastlyServicesCmd)
}
