package fleet_automation

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteFleetScheduleCmd = &cobra.Command{
	Use:   "deletefleetschedule [id]",
	Short: "Delete a schedule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		_, err := api.DeleteFleetSchedule(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deletefleetschedule: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteFleetScheduleCmd)
}
