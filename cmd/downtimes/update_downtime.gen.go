package downtimes

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateDowntimeCmd = &cobra.Command{
	Use:     "update-downtime [downtime_id]",
	Aliases: []string{"update"},
	Short:   "Update a downtime",
	Long: `Update a downtime
Documentation: https://docs.datadoghq.com/api/latest/downtimes/#update-downtime`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DowntimeResponse
		var err error

		var body datadogV2.DowntimeUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewDowntimesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateDowntime(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-downtime")

		cmd.Println(cmdutil.FormatJSON(res, "downtime"))
	},
}

func init() {

	UpdateDowntimeCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateDowntimeCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateDowntimeCmd)
}
