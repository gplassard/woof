package apm

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetServiceListCmd = &cobra.Command{
	Use:   "get-service-list",
	
	Short: "Get service list",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAPMApi(client.NewAPIClient())
		res, _, err := api.GetServiceList(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to get-service-list: %v", err)
		}

		cmdutil.PrintJSON(res, "services_list")
	},
}

func init() {
	Cmd.AddCommand(GetServiceListCmd)
}
