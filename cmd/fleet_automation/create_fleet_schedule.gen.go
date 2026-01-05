package fleet_automation

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateFleetScheduleCmd = &cobra.Command{
	Use: "create-fleet-schedule",

	Short: "Create a schedule",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err := api.CreateFleetSchedule(client.NewContext(apiKey, appKey, site), datadogV2.FleetScheduleCreateRequest{})
		cmdutil.HandleError(err, "failed to create-fleet-schedule")

		cmd.Println(cmdutil.FormatJSON(res, "schedule"))
	},
}

func init() {
	Cmd.AddCommand(CreateFleetScheduleCmd)
}
