package fleet_automation

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateFleetScheduleCmd = &cobra.Command{
	Use:   "create-fleet-schedule",
	
	Short: "Create a schedule",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err := api.CreateFleetSchedule(client.NewContext(apiKey, appKey, site), datadogV2.FleetScheduleCreateRequest{})
		cmdutil.HandleError(err, "failed to create-fleet-schedule")

		cmdutil.PrintJSON(res, "schedule")
	},
}

func init() {
	Cmd.AddCommand(CreateFleetScheduleCmd)
}
