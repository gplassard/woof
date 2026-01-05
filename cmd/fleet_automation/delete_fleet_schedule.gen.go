package fleet_automation

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteFleetScheduleCmd = &cobra.Command{
	Use: "delete-fleet-schedule [id]",

	Short: "Delete a schedule",
	Long: `Delete a schedule
Documentation: https://docs.datadoghq.com/api/latest/fleet-automation/#delete-fleet-schedule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		_, err = api.DeleteFleetSchedule(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-fleet-schedule")

	},
}

func init() {
	Cmd.AddCommand(DeleteFleetScheduleCmd)
}
