package fleet_automation

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteFleetScheduleCmd = &cobra.Command{
	Use:   "delete_fleet_schedule [id]",
	Short: "Delete a schedule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		_, err := api.DeleteFleetSchedule(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete_fleet_schedule: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteFleetScheduleCmd)
}
