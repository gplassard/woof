package actions_datastores

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListDatastoresCmd = &cobra.Command{
	Use:   "list-datastores",
	Aliases: []string{ "list", },
	Short: "List datastores",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		res, _, err := api.ListDatastores(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-datastores")

		cmdutil.PrintJSON(res, "datastores")
	},
}

func init() {
	Cmd.AddCommand(ListDatastoresCmd)
}
