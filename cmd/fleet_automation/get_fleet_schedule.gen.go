package fleet_automation

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetFleetScheduleCmd = &cobra.Command{
	Use:   "get-fleet-schedule [id]",
	
	Short: "Get a schedule by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err := api.GetFleetSchedule(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-fleet-schedule")

		cmd.Println(cmdutil.FormatJSON(res, "schedule"))
	},
}

func init() {
	Cmd.AddCommand(GetFleetScheduleCmd)
}
