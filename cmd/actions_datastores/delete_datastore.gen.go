package actions_datastores

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteDatastoreCmd = &cobra.Command{
	Use:     "delete-datastore [datastore_id]",
	Aliases: []string{"delete"},
	Short:   "Delete datastore",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		_, err := api.DeleteDatastore(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-datastore")

	},
}

func init() {
	Cmd.AddCommand(DeleteDatastoreCmd)
}
