package fleet_automation

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateFleetScheduleCmd = &cobra.Command{
	Use: "create-fleet-schedule [payload]",

	Short: "Create a schedule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.FleetScheduleCreateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err := api.CreateFleetSchedule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-fleet-schedule")

		cmd.Println(cmdutil.FormatJSON(res, "schedule"))
	},
}

func init() {
	Cmd.AddCommand(CreateFleetScheduleCmd)
}
