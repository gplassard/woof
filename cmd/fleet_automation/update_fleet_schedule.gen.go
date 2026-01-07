package fleet_automation

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateFleetScheduleCmd = &cobra.Command{
	Use: "update-fleet-schedule [id]",

	Short: "Update a schedule",
	Long: `Update a schedule
Documentation: https://docs.datadoghq.com/api/latest/fleet-automation/#update-fleet-schedule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FleetScheduleResponse
		var err error

		var body datadogV2.FleetSchedulePatchRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateFleetSchedule(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-fleet-schedule")

		fmt.Println(cmdutil.FormatJSON(res, "schedule"))
	},
}

func init() {

	UpdateFleetScheduleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateFleetScheduleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateFleetScheduleCmd)
}
