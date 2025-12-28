package fleet_automation

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateFleetScheduleCmd = &cobra.Command{
	Use:   "update_fleet_schedule [id]",
	Short: "Update a schedule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err := api.UpdateFleetSchedule(client.NewContext(apiKey, appKey, site), args[0], datadogV2.FleetSchedulePatchRequest{})
		if err != nil {
			log.Fatalf("failed to update_fleet_schedule: %v", err)
		}

		cmdutil.PrintJSON(res, "schedule")
	},
}

func init() {
	Cmd.AddCommand(UpdateFleetScheduleCmd)
}
