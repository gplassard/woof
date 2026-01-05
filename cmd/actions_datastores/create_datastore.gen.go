package actions_datastores

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateDatastoreCmd = &cobra.Command{
	Use:     "create-datastore [payload]",
	Aliases: []string{"create"},
	Short:   "Create datastore",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CreateAppsDatastoreResponse
		var err error

		var body datadogV2.CreateAppsDatastoreRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		res, _, err = api.CreateDatastore(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-datastore")

		cmd.Println(cmdutil.FormatJSON(res, "datastores"))
	},
}

func init() {
	Cmd.AddCommand(CreateDatastoreCmd)
}
