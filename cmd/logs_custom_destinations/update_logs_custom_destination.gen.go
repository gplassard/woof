package logs_custom_destinations

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateLogsCustomDestinationCmd = &cobra.Command{
	Use:     "update-logs-custom-destination [custom_destination_id] [payload]",
	Aliases: []string{"update"},
	Short:   "Update a custom destination",
	Long: `Update a custom destination
Documentation: https://docs.datadoghq.com/api/latest/logs-custom-destinations/#update-logs-custom-destination`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CustomDestinationResponse
		var err error

		var body datadogV2.CustomDestinationUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewLogsCustomDestinationsApi(client.NewAPIClient())
		res, _, err = api.UpdateLogsCustomDestination(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-logs-custom-destination")

		cmd.Println(cmdutil.FormatJSON(res, "custom_destination"))
	},
}

func init() {
	Cmd.AddCommand(UpdateLogsCustomDestinationCmd)
}
