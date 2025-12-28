package on_call

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetScheduleOnCallUserCmd = &cobra.Command{
	Use:   "get-schedule-on-call-user [schedule_id]",
	Aliases: []string{ "get-schedule-user", },
	Short: "Get scheduled on-call user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.GetScheduleOnCallUser(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-schedule-on-call-user")

		cmdutil.PrintJSON(res, "shifts")
	},
}

func init() {
	Cmd.AddCommand(GetScheduleOnCallUserCmd)
}
