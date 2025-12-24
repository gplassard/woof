package actions_datastores

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateDatastoreCmd = &cobra.Command{
	Use:   "createdatastore",
	Short: "Create datastore",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		res, _, err := api.CreateDatastore(client.NewContext(apiKey, appKey, site), datadogV2.CreateAppsDatastoreRequest{})
		if err != nil {
			log.Fatalf("failed to createdatastore: %v", err)
		}

		cmdutil.PrintJSON(res, "actions_datastores")
	},
}

func init() {
	Cmd.AddCommand(CreateDatastoreCmd)
}
