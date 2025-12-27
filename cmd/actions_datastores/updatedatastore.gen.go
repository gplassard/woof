package actions_datastores

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateDatastoreCmd = &cobra.Command{
	Use:   "updatedatastore [datastore_id]",
	Short: "Update datastore",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		res, _, err := api.UpdateDatastore(client.NewContext(apiKey, appKey, site), args[0], datadogV2.UpdateAppsDatastoreRequest{})
		if err != nil {
			log.Fatalf("failed to updatedatastore: %v", err)
		}

		cmdutil.PrintJSON(res, "actions_datastores")
	},
}

func init() {
	Cmd.AddCommand(UpdateDatastoreCmd)
}
