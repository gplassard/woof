package on_call

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateOnCallScheduleCmd = &cobra.Command{
	Use:   "update-on-call-schedule [schedule_id]",
	Aliases: []string{ "update-schedule", },
	Short: "Update On-Call schedule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.UpdateOnCallSchedule(client.NewContext(apiKey, appKey, site), args[0], datadogV2.ScheduleUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-on-call-schedule")

		cmdutil.PrintJSON(res, "schedules")
	},
}

func init() {
	Cmd.AddCommand(UpdateOnCallScheduleCmd)
}
