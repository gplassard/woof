package fleet_automation

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListFleetSchedulesCmd = &cobra.Command{
	Use: "list-fleet-schedules",

	Short: "List all schedules",
	Long: `List all schedules
Documentation: https://docs.datadoghq.com/api/latest/fleet-automation/#list-fleet-schedules`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FleetSchedulesResponse
		var err error

		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err = api.ListFleetSchedules(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-fleet-schedules")

		cmd.Println(cmdutil.FormatJSON(res, "schedule"))
	},
}

func init() {
	Cmd.AddCommand(ListFleetSchedulesCmd)
}
