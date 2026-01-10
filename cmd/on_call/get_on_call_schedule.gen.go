package on_call

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetOnCallScheduleCmd = &cobra.Command{
	Use:     "get-on-call-schedule [schedule_id]",
	Aliases: []string{"get-schedule"},
	Short:   "Get On-Call schedule",
	Long: `Get On-Call schedule
Documentation: https://docs.datadoghq.com/api/latest/on-call/#get-on-call-schedule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.Schedule
		var err error

		optionalParams := datadogV2.NewGetOnCallScheduleOptionalParameters()

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetOnCallSchedule(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to get-on-call-schedule")

		fmt.Println(cmdutil.FormatJSON(res, "schedules"))
	},
}

func init() {

	GetOnCallScheduleCmd.Flags().String("include", "", "Comma-separated list of included relationships to be returned. Allowed values: 'teams', 'layers', 'layers.members', 'layers.members.user'.")

	Cmd.AddCommand(GetOnCallScheduleCmd)
}
