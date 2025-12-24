package app_builder

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteAppsCmd = &cobra.Command{
	Use:   "deleteapps",
	Short: "Delete Multiple Apps",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewAppBuilderApi(client.NewAPIClient())
		res, _, err := api.DeleteApps(client.NewContext(apiKey, appKey, site), datadogV2.DeleteAppsRequest{})
		if err != nil {
			log.Fatalf("failed to deleteapps: %v", err)
		}

		cmdutil.PrintJSON(res, "app_builder")
	},
}

func init() {
	Cmd.AddCommand(DeleteAppsCmd)
}
