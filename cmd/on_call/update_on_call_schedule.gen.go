package on_call

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateOnCallScheduleCmd = &cobra.Command{
	Use:     "update-on-call-schedule [schedule_id] [payload]",
	Aliases: []string{"update-schedule"},
	Short:   "Update On-Call schedule",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.Schedule
		var err error

		var body datadogV2.ScheduleUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err = api.UpdateOnCallSchedule(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-on-call-schedule")

		cmd.Println(cmdutil.FormatJSON(res, "schedules"))
	},
}

func init() {
	Cmd.AddCommand(UpdateOnCallScheduleCmd)
}
