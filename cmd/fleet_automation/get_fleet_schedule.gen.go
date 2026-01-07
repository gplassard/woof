package fleet_automation

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetFleetScheduleCmd = &cobra.Command{
	Use: "get-fleet-schedule [id]",

	Short: "Get a schedule by ID",
	Long: `Get a schedule by ID
Documentation: https://docs.datadoghq.com/api/latest/fleet-automation/#get-fleet-schedule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FleetScheduleResponse
		var err error

		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetFleetSchedule(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-fleet-schedule")

		fmt.Println(cmdutil.FormatJSON(res, "schedule"))
	},
}

func init() {

	Cmd.AddCommand(GetFleetScheduleCmd)
}
