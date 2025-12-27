package downtimes

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListDowntimesCmd = &cobra.Command{
	Use:   "list_downtimes",
	Short: "Get all downtimes",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDowntimesApi(client.NewAPIClient())
		res, _, err := api.ListDowntimes(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list_downtimes: %v", err)
		}

		cmdutil.PrintJSON(res, "downtimes")
	},
}

func init() {
	Cmd.AddCommand(ListDowntimesCmd)
}
