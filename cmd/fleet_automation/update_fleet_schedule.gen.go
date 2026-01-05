package fleet_automation

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateFleetScheduleCmd = &cobra.Command{
	Use: "update-fleet-schedule [id]",

	Short: "Update a schedule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err := api.UpdateFleetSchedule(client.NewContext(apiKey, appKey, site), args[0], datadogV2.FleetSchedulePatchRequest{})
		cmdutil.HandleError(err, "failed to update-fleet-schedule")

		cmd.Println(cmdutil.FormatJSON(res, "schedule"))
	},
}

func init() {
	Cmd.AddCommand(UpdateFleetScheduleCmd)
}
