package on_call

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateOnCallScheduleCmd = &cobra.Command{
	Use:     "update-on-call-schedule [schedule_id]",
	Aliases: []string{"update-schedule"},
	Short:   "Update On-Call schedule",
	Long: `Update On-Call schedule
Documentation: https://docs.datadoghq.com/api/latest/on-call/#update-on-call-schedule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.Schedule
		var err error

		optionalParams := datadogV2.NewUpdateOnCallScheduleOptionalParameters()

		if cmd.Flags().Changed("payload") || cmd.Flags().Changed("payload-file") {
			err = cmdutil.UnmarshalPayload(cmd, optionalParams)
			cmdutil.HandleError(err, "failed to read payload")
		}

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateOnCallSchedule(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to update-on-call-schedule")

		fmt.Println(cmdutil.FormatJSON(res, "schedules"))
	},
}

func init() {

	UpdateOnCallScheduleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateOnCallScheduleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	UpdateOnCallScheduleCmd.Flags().String("include", "", "Comma-separated list of included relationships to be returned. Allowed values: 'teams', 'layers', 'layers.members', 'layers.members.user'.")

	Cmd.AddCommand(UpdateOnCallScheduleCmd)
}
