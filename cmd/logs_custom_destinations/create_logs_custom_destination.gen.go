package logs_custom_destinations

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateLogsCustomDestinationCmd = &cobra.Command{
	Use:   "create_logs_custom_destination",
	Short: "Create a custom destination",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsCustomDestinationsApi(client.NewAPIClient())
		res, _, err := api.CreateLogsCustomDestination(client.NewContext(apiKey, appKey, site), datadogV2.CustomDestinationCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create_logs_custom_destination: %v", err)
		}

		cmdutil.PrintJSON(res, "logs_custom_destinations")
	},
}

func init() {
	Cmd.AddCommand(CreateLogsCustomDestinationCmd)
}
