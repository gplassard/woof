package actions_datastores

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateDatastoreItemCmd = &cobra.Command{
	Use:     "update-datastore-item [datastore_id]",
	Aliases: []string{"update-item"},
	Short:   "Update datastore item",
	Long: `Update datastore item
Documentation: https://docs.datadoghq.com/api/latest/actions-datastores/#update-datastore-item`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ItemApiPayload
		var err error

		var body datadogV2.UpdateAppsDatastoreItemRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateDatastoreItem(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-datastore-item")

		cmd.Println(cmdutil.FormatJSON(res, "items"))
	},
}

func init() {

	UpdateDatastoreItemCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateDatastoreItemCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateDatastoreItemCmd)
}
