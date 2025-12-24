package actions_datastores

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateDatastoreItemCmd = &cobra.Command{
	Use:   "updatedatastoreitem [datastore_id]",
	Short: "Update datastore item",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		res, _, err := api.UpdateDatastoreItem(client.NewContext(apiKey, appKey, site), args[0], datadogV2.UpdateAppsDatastoreItemRequest{})
		if err != nil {
			log.Fatalf("failed to updatedatastoreitem: %v", err)
		}

		cmdutil.PrintJSON(res, "actions_datastores")
	},
}

func init() {
	Cmd.AddCommand(UpdateDatastoreItemCmd)
}
