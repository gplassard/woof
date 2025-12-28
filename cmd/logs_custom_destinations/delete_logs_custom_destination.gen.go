package logs_custom_destinations

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteLogsCustomDestinationCmd = &cobra.Command{
	Use:   "delete-logs-custom-destination [custom_destination_id]",
	Short: "Delete a custom destination",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsCustomDestinationsApi(client.NewAPIClient())
		_, err := api.DeleteLogsCustomDestination(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete-logs-custom-destination: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteLogsCustomDestinationCmd)
}
