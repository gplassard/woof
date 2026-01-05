package on_call

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateOnCallScheduleCmd = &cobra.Command{
	Use:     "create-on-call-schedule",
	Aliases: []string{"create-schedule"},
	Short:   "Create On-Call schedule",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.CreateOnCallSchedule(client.NewContext(apiKey, appKey, site), datadogV2.ScheduleCreateRequest{})
		cmdutil.HandleError(err, "failed to create-on-call-schedule")

		cmd.Println(cmdutil.FormatJSON(res, "schedules"))
	},
}

func init() {
	Cmd.AddCommand(CreateOnCallScheduleCmd)
}
