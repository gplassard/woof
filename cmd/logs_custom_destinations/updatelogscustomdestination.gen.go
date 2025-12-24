package logs_custom_destinations

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateLogsCustomDestinationCmd = &cobra.Command{
	Use:   "updatelogscustomdestination [custom_destination_id]",
	Short: "Update a custom destination",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewLogsCustomDestinationsApi(client.NewAPIClient())
		res, _, err := api.UpdateLogsCustomDestination(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CustomDestinationUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to updatelogscustomdestination: %v", err)
		}

		cmdutil.PrintJSON(res, "logs_custom_destinations")
	},
}

func init() {
	Cmd.AddCommand(UpdateLogsCustomDestinationCmd)
}
