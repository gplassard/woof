package actions_datastores

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetDatastoreCmd = &cobra.Command{
	Use:     "get-datastore [datastore_id]",
	Aliases: []string{"get"},
	Short:   "Get datastore",
	Long: `Get datastore
Documentation: https://docs.datadoghq.com/api/latest/actions-datastores/#get-datastore`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.Datastore
		var err error

		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		res, _, err = api.GetDatastore(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-datastore")

		cmd.Println(cmdutil.FormatJSON(res, "datastores"))
	},
}

func init() {
	Cmd.AddCommand(GetDatastoreCmd)
}
