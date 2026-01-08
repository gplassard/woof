package api_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateOpenAPICmd = &cobra.Command{
	Use: "create-open-api",

	Short: "Create a new API",
	Long: `Create a new API
Documentation: https://docs.datadoghq.com/api/latest/a-p-i-management/#create-open-api`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CreateOpenAPIResponse
		var err error

		api := datadogV2.NewAPIManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateOpenAPI(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to create-open-api")

		fmt.Println(cmdutil.FormatJSON(res, "open_a_p_i"))
	},
}

func init() {

	Cmd.AddCommand(CreateOpenAPICmd)
}
