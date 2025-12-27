package fleet_automation

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateFleetScheduleCmd = &cobra.Command{
	Use:   "create_fleet_schedule",
	Short: "Create a schedule",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err := api.CreateFleetSchedule(client.NewContext(apiKey, appKey, site), datadogV2.FleetScheduleCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create_fleet_schedule: %v", err)
		}

		cmdutil.PrintJSON(res, "fleet_automation")
	},
}

func init() {
	Cmd.AddCommand(CreateFleetScheduleCmd)
}
