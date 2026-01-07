package fleet_automation

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateFleetScheduleCmd = &cobra.Command{
	Use: "create-fleet-schedule",

	Short: "Create a schedule",
	Long: `Create a schedule
Documentation: https://docs.datadoghq.com/api/latest/fleet-automation/#create-fleet-schedule`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FleetScheduleResponse
		var err error

		var body datadogV2.FleetScheduleCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateFleetSchedule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-fleet-schedule")

		fmt.Println(cmdutil.FormatJSON(res, "schedule"))
	},
}

func init() {

	CreateFleetScheduleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateFleetScheduleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateFleetScheduleCmd)
}
