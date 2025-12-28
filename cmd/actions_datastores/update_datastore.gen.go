package actions_datastores

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateDatastoreCmd = &cobra.Command{
	Use:   "update-datastore [datastore_id]",
	Aliases: []string{ "update", },
	Short: "Update datastore",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		res, _, err := api.UpdateDatastore(client.NewContext(apiKey, appKey, site), args[0], datadogV2.UpdateAppsDatastoreRequest{})
		cmdutil.HandleError(err, "failed to update-datastore")

		cmdutil.PrintJSON(res, "datastores")
	},
}

func init() {
	Cmd.AddCommand(UpdateDatastoreCmd)
}
