package app_builder

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateAppCmd = &cobra.Command{
	Use: "update-app [app_id] [payload]",

	Short: "Update App",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.UpdateAppRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewAppBuilderApi(client.NewAPIClient())
		res, _, err := api.UpdateApp(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), body)
		cmdutil.HandleError(err, "failed to update-app")

		cmd.Println(cmdutil.FormatJSON(res, "appDefinitions"))
	},
}

func init() {
	Cmd.AddCommand(UpdateAppCmd)
}
