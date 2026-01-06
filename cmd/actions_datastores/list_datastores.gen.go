package actions_datastores

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListDatastoresCmd = &cobra.Command{
	Use:     "list-datastores",
	Aliases: []string{"list"},
	Short:   "List datastores",
	Long: `List datastores
Documentation: https://docs.datadoghq.com/api/latest/actions-datastores/#list-datastores`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DatastoreArray
		var err error

		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListDatastores(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-datastores")

		cmd.Println(cmdutil.FormatJSON(res, "datastores"))
	},
}

func init() {

	Cmd.AddCommand(ListDatastoresCmd)
}
