package rum

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetRUMApplicationCmd = &cobra.Command{
	Use:   "get-rum-application [id]",
	Aliases: []string{ "get-application", },
	Short: "Get a RUM application",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRUMApi(client.NewAPIClient())
		res, _, err := api.GetRUMApplication(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-rum-application")

		cmdutil.PrintJSON(res, "rum_application")
	},
}

func init() {
	Cmd.AddCommand(GetRUMApplicationCmd)
}
