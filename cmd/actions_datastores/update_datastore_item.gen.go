package actions_datastores

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateDatastoreItemCmd = &cobra.Command{
	Use:     "update-datastore-item [datastore_id] [payload]",
	Aliases: []string{"update-item"},
	Short:   "Update datastore item",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ItemApiPayload
		var err error

		var body datadogV2.UpdateAppsDatastoreItemRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		res, _, err = api.UpdateDatastoreItem(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-datastore-item")

		cmd.Println(cmdutil.FormatJSON(res, "items"))
	},
}

func init() {
	Cmd.AddCommand(UpdateDatastoreItemCmd)
}
