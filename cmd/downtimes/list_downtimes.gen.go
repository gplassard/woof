package downtimes

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListDowntimesCmd = &cobra.Command{
	Use:   "list-downtimes",
	Aliases: []string{ "list", },
	Short: "Get all downtimes",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDowntimesApi(client.NewAPIClient())
		res, _, err := api.ListDowntimes(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-downtimes")

		cmdutil.PrintJSON(res, "downtime")
	},
}

func init() {
	Cmd.AddCommand(ListDowntimesCmd)
}
