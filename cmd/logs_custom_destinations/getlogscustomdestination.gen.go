package logs_custom_destinations

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetLogsCustomDestinationCmd = &cobra.Command{
	Use:   "getlogscustomdestination [custom_destination_id]",
	Short: "Get a custom destination",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewLogsCustomDestinationsApi(client.NewAPIClient())
		res, _, err := api.GetLogsCustomDestination(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getlogscustomdestination: %v", err)
		}

		cmdutil.PrintJSON(res, "logs_custom_destinations")
	},
}

func init() {
	Cmd.AddCommand(GetLogsCustomDestinationCmd)
}
