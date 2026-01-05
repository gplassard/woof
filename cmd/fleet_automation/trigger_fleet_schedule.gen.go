package fleet_automation

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var TriggerFleetScheduleCmd = &cobra.Command{
	Use: "trigger-fleet-schedule [id]",

	Short: "Trigger a schedule deployment",
	Long: `Trigger a schedule deployment
Documentation: https://docs.datadoghq.com/api/latest/fleet-automation/#trigger-fleet-schedule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FleetDeploymentResponse
		var err error

		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err = api.TriggerFleetSchedule(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to trigger-fleet-schedule")

		cmd.Println(cmdutil.FormatJSON(res, "deployment"))
	},
}

func init() {

	Cmd.AddCommand(TriggerFleetScheduleCmd)
}
