package actions_datastores

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateDatastoreCmd = &cobra.Command{
	Use:     "update-datastore [datastore_id] [payload]",
	Aliases: []string{"update"},
	Short:   "Update datastore",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.UpdateAppsDatastoreRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		res, _, err := api.UpdateDatastore(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-datastore")

		cmd.Println(cmdutil.FormatJSON(res, "datastores"))
	},
}

func init() {
	Cmd.AddCommand(UpdateDatastoreCmd)
}
