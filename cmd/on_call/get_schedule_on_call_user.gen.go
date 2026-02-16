package on_call

import (
	"fmt"
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

		optionalParams := datadogV2.NewGetScheduleOnCallUserOptionalParameters()

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		if cmd.Flags().Changed("filter-at-ts") {
			val, _ := cmd.Flags().GetString("filter-at-ts")
			optionalParams.WithFilterAtTs(val)
		}

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetScheduleOnCallUser(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to get-schedule-on-call-user")

		fmt.Println(cmdutil.FormatJSON(res, "shifts"))
	},
}

func init() {

	GetScheduleOnCallUserCmd.Flags().String("include", "", "Specifies related resources to include in the response as a comma-separated list. Allowed value: 'user'.")

	GetScheduleOnCallUserCmd.Flags().String("filter-at-ts", "", "Retrieves the on-call user at the given timestamp (ISO-8601). Defaults to the current time if omitted.\"")

	Cmd.AddCommand(GetScheduleOnCallUserCmd)
}
