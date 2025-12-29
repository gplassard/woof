package fleet_automation

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListFleetSchedulesCmd = &cobra.Command{
	Use:   "list-fleet-schedules",
	
	Short: "List all schedules",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err := api.ListFleetSchedules(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-fleet-schedules")

		cmd.Println(cmdutil.FormatJSON(res, "schedule"))
	},
}

func init() {
	Cmd.AddCommand(ListFleetSchedulesCmd)
}
