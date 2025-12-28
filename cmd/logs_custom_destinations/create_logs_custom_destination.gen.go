package logs_custom_destinations

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateLogsCustomDestinationCmd = &cobra.Command{
	Use:   "create-logs-custom-destination",
	Aliases: []string{ "create", },
	Short: "Create a custom destination",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsCustomDestinationsApi(client.NewAPIClient())
		res, _, err := api.CreateLogsCustomDestination(client.NewContext(apiKey, appKey, site), datadogV2.CustomDestinationCreateRequest{})
		cmdutil.HandleError(err, "failed to create-logs-custom-destination")

		cmdutil.PrintJSON(res, "custom_destination")
	},
}

func init() {
	Cmd.AddCommand(CreateLogsCustomDestinationCmd)
}
