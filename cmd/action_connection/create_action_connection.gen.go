package action_connection

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateActionConnectionCmd = &cobra.Command{
	Use:     "create-action-connection [payload]",
	Aliases: []string{"create"},
	Short:   "Create a new Action Connection",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CreateActionConnectionResponse
		var err error

		var body datadogV2.CreateActionConnectionRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewActionConnectionApi(client.NewAPIClient())
		res, _, err = api.CreateActionConnection(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-action-connection")

		cmd.Println(cmdutil.FormatJSON(res, "action_connection"))
	},
}

func init() {
	Cmd.AddCommand(CreateActionConnectionCmd)
}
