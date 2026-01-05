package on_call

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateOnCallScheduleCmd = &cobra.Command{
	Use:     "create-on-call-schedule [payload]",
	Aliases: []string{"create-schedule"},
	Short:   "Create On-Call schedule",
	Long: `Create On-Call schedule
Documentation: https://docs.datadoghq.com/api/latest/on-call/#create-on-call-schedule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.Schedule
		var err error

		var body datadogV2.ScheduleCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err = api.CreateOnCallSchedule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-on-call-schedule")

		cmd.Println(cmdutil.FormatJSON(res, "schedules"))
	},
}

func init() {
	Cmd.AddCommand(CreateOnCallScheduleCmd)
}
