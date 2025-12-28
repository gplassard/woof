package logs_custom_destinations

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetLogsCustomDestinationCmd = &cobra.Command{
	Use:   "get-logs-custom-destination [custom_destination_id]",
	Aliases: []string{ "get", },
	Short: "Get a custom destination",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsCustomDestinationsApi(client.NewAPIClient())
		res, _, err := api.GetLogsCustomDestination(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-logs-custom-destination")

		cmdutil.PrintJSON(res, "custom_destination")
	},
}

func init() {
	Cmd.AddCommand(GetLogsCustomDestinationCmd)
}
