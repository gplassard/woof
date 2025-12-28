package logs_custom_destinations

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateLogsCustomDestinationCmd = &cobra.Command{
	Use:   "update-logs-custom-destination [custom_destination_id]",
	Aliases: []string{ "update", },
	Short: "Update a custom destination",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsCustomDestinationsApi(client.NewAPIClient())
		res, _, err := api.UpdateLogsCustomDestination(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CustomDestinationUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-logs-custom-destination")

		cmdutil.PrintJSON(res, "custom_destination")
	},
}

func init() {
	Cmd.AddCommand(UpdateLogsCustomDestinationCmd)
}
