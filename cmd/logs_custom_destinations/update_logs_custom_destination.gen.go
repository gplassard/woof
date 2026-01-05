package logs_custom_destinations

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateLogsCustomDestinationCmd = &cobra.Command{
	Use:     "update-logs-custom-destination [custom_destination_id]",
	Aliases: []string{"update"},
	Short:   "Update a custom destination",
	Long: `Update a custom destination
Documentation: https://docs.datadoghq.com/api/latest/logs-custom-destinations/#update-logs-custom-destination`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CustomDestinationResponse
		var err error

		var body datadogV2.CustomDestinationUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewLogsCustomDestinationsApi(client.NewAPIClient())
		res, _, err = api.UpdateLogsCustomDestination(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-logs-custom-destination")

		cmd.Println(cmdutil.FormatJSON(res, "custom_destination"))
	},
}

func init() {

	UpdateLogsCustomDestinationCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateLogsCustomDestinationCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateLogsCustomDestinationCmd)
}
