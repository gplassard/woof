package action_connection

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateActionConnectionCmd = &cobra.Command{
	Use:     "update-action-connection [connection_id] [payload]",
	Aliases: []string{"update"},
	Short:   "Update an existing Action Connection",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.UpdateActionConnectionRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewActionConnectionApi(client.NewAPIClient())
		res, _, err := api.UpdateActionConnection(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-action-connection")

		cmd.Println(cmdutil.FormatJSON(res, "action_connection"))
	},
}

func init() {
	Cmd.AddCommand(UpdateActionConnectionCmd)
}
