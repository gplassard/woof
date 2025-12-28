package downtimes

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CancelDowntimeCmd = &cobra.Command{
	Use:   "cancel-downtime [downtime_id]",
	Short: "Cancel a downtime",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDowntimesApi(client.NewAPIClient())
		_, err := api.CancelDowntime(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to cancel-downtime: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(CancelDowntimeCmd)
}
