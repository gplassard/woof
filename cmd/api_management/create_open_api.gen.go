package api_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateOpenAPICmd = &cobra.Command{
	Use: "create-open-api [payload]",

	Short: "Create a new API",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CreateOpenAPIResponse
		var err error

		var body datadogV2.CreateOpenAPIOptionalParameters
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewAPIManagementApi(client.NewAPIClient())
		res, _, err = api.CreateOpenAPI(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-open-api")

		cmd.Println(cmdutil.FormatJSON(res, "api_management"))
	},
}

func init() {
	Cmd.AddCommand(CreateOpenAPICmd)
}
