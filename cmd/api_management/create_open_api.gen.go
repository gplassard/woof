package api_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateOpenAPICmd = &cobra.Command{
	Use: "create-open-api",

	Short: "Create a new API",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAPIManagementApi(client.NewAPIClient())
		res, _, err := api.CreateOpenAPI(client.NewContext(apiKey, appKey, site), *datadogV2.NewCreateOpenAPIOptionalParameters())
		cmdutil.HandleError(err, "failed to create-open-api")

		cmd.Println(cmdutil.FormatJSON(res, "api_management"))
	},
}

func init() {
	Cmd.AddCommand(CreateOpenAPICmd)
}
