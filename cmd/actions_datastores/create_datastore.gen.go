package actions_datastores

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateDatastoreCmd = &cobra.Command{
	Use:   "create-datastore",
	Aliases: []string{ "create", },
	Short: "Create datastore",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		res, _, err := api.CreateDatastore(client.NewContext(apiKey, appKey, site), datadogV2.CreateAppsDatastoreRequest{})
		cmdutil.HandleError(err, "failed to create-datastore")

		cmdutil.PrintJSON(res, "datastores")
	},
}

func init() {
	Cmd.AddCommand(CreateDatastoreCmd)
}
