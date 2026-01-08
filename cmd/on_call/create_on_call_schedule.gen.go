package on_call

import (
	"fmt"
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
	Long: `Create On-Call schedule
Documentation: https://docs.datadoghq.com/api/latest/on-call/#create-on-call-schedule`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.Schedule
		var err error

		var body datadogV2.ScheduleCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateOnCallSchedule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-on-call-schedule")

		fmt.Println(cmdutil.FormatJSON(res, "on_call_schedule"))
	},
}

func init() {

	CreateOnCallScheduleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateOnCallScheduleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateOnCallScheduleCmd)
}
