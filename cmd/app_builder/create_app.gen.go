package app_builder

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateAppCmd = &cobra.Command{
	Use: "create-app [payload]",

	Short: "Create App",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CreateAppResponse
		var err error

		var body datadogV2.CreateAppRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewAppBuilderApi(client.NewAPIClient())
		res, _, err = api.CreateApp(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-app")

		cmd.Println(cmdutil.FormatJSON(res, "appDefinitions"))
	},
}

func init() {
	Cmd.AddCommand(CreateAppCmd)
}
