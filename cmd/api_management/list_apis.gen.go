package api_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAPIsCmd = &cobra.Command{
	Use: "list-apis",

	Short: "List APIs",
	Long: `List APIs
Documentation: https://docs.datadoghq.com/api/latest/api-management/#list-apis`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListAPIsResponse
		var err error

		api := datadogV2.NewAPIManagementApi(client.NewAPIClient())
		res, _, err = api.ListAPIs(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-apis")

		cmd.Println(cmdutil.FormatJSON(res, "api_management"))
	},
}

func init() {
	Cmd.AddCommand(ListAPIsCmd)
}
