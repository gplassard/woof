package fastly_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateFastlyServiceCmd = &cobra.Command{
	Use:   "update-fastly-service [account_id] [service_id]",
	Short: "Update Fastly service",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		res, _, err := api.UpdateFastlyService(client.NewContext(apiKey, appKey, site), args[0], args[1], datadogV2.FastlyServiceRequest{})
		if err != nil {
			log.Fatalf("failed to update-fastly-service: %v", err)
		}

		cmdutil.PrintJSON(res, "fastly-services")
	},
}

func init() {
	Cmd.AddCommand(UpdateFastlyServiceCmd)
}
