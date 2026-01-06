package on_call

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetScheduleOnCallUserCmd = &cobra.Command{
	Use:     "get-schedule-on-call-user [schedule_id]",
	Aliases: []string{"get-schedule-user"},
	Short:   "Get scheduled on-call user",
	Long: `Get scheduled on-call user
Documentation: https://docs.datadoghq.com/api/latest/on-call/#get-schedule-on-call-user`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.Shift
		var err error

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetScheduleOnCallUser(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-schedule-on-call-user")

		cmd.Println(cmdutil.FormatJSON(res, "shifts"))
	},
}

func init() {

	Cmd.AddCommand(GetScheduleOnCallUserCmd)
}
