package logs_custom_destinations

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListLogsCustomDestinationsCmd = &cobra.Command{
	Use:   "list-logs-custom-destinations",
	Short: "Get all custom destinations",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsCustomDestinationsApi(client.NewAPIClient())
		res, _, err := api.ListLogsCustomDestinations(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-logs-custom-destinations: %v", err)
		}

		cmdutil.PrintJSON(res, "custom_destination")
	},
}

func init() {
	Cmd.AddCommand(ListLogsCustomDestinationsCmd)
}
