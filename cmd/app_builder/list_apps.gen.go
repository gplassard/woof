package app_builder

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListAppsCmd = &cobra.Command{
	Use:   "list-apps",
	
	Short: "List Apps",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAppBuilderApi(client.NewAPIClient())
		res, _, err := api.ListApps(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-apps")

		cmdutil.PrintJSON(res, "appDefinitions")
	},
}

func init() {
	Cmd.AddCommand(ListAppsCmd)
}
