package logs_custom_destinations

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetLogsCustomDestinationCmd = &cobra.Command{
	Use:     "get-logs-custom-destination [custom_destination_id]",
	Aliases: []string{"get"},
	Short:   "Get a custom destination",
	Long: `Get a custom destination
Documentation: https://docs.datadoghq.com/api/latest/logs-custom-destinations/#get-logs-custom-destination`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CustomDestinationResponse
		var err error

		api := datadogV2.NewLogsCustomDestinationsApi(client.NewAPIClient())
		res, _, err = api.GetLogsCustomDestination(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-logs-custom-destination")

		cmd.Println(cmdutil.FormatJSON(res, "custom_destination"))
	},
}

func init() {
	Cmd.AddCommand(GetLogsCustomDestinationCmd)
}
