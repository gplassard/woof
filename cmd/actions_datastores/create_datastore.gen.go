package actions_datastores

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateDatastoreCmd = &cobra.Command{
	Use:     "create-datastore",
	Aliases: []string{"create"},
	Short:   "Create datastore",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		res, _, err := api.CreateDatastore(client.NewContext(apiKey, appKey, site), datadogV2.CreateAppsDatastoreRequest{})
		cmdutil.HandleError(err, "failed to create-datastore")

		cmd.Println(cmdutil.FormatJSON(res, "datastores"))
	},
}

func init() {
	Cmd.AddCommand(CreateDatastoreCmd)
}
