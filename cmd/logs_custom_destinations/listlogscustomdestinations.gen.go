package logs_custom_destinations

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListLogsCustomDestinationsCmd = &cobra.Command{
	Use:   "listlogscustomdestinations",
	Short: "Get all custom destinations",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewLogsCustomDestinationsApi(client.NewAPIClient())
		res, _, err := api.ListLogsCustomDestinations(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listlogscustomdestinations: %v", err)
		}

		cmdutil.PrintJSON(res, "logs_custom_destinations")
	},
}

func init() {
	Cmd.AddCommand(ListLogsCustomDestinationsCmd)
}
