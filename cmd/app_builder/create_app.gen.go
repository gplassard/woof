package app_builder

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateAppCmd = &cobra.Command{
	Use:   "create-app",
	
	Short: "Create App",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAppBuilderApi(client.NewAPIClient())
		res, _, err := api.CreateApp(client.NewContext(apiKey, appKey, site), datadogV2.CreateAppRequest{})
		cmdutil.HandleError(err, "failed to create-app")

		cmdutil.PrintJSON(res, "appDefinitions")
	},
}

func init() {
	Cmd.AddCommand(CreateAppCmd)
}
