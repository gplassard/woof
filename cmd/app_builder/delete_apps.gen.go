package app_builder

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteAppsCmd = &cobra.Command{
	Use: "delete-apps",

	Short: "Delete Multiple Apps",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAppBuilderApi(client.NewAPIClient())
		res, _, err := api.DeleteApps(client.NewContext(apiKey, appKey, site), datadogV2.DeleteAppsRequest{})
		cmdutil.HandleError(err, "failed to delete-apps")

		cmd.Println(cmdutil.FormatJSON(res, "appDefinitions"))
	},
}

func init() {
	Cmd.AddCommand(DeleteAppsCmd)
}
