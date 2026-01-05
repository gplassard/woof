package api_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var GetOpenAPICmd = &cobra.Command{
	Use: "get-open-api [id]",

	Short: "Get an API",
	Long: `Get an API
Documentation: https://docs.datadoghq.com/api/latest/api-management/#get-open-api`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res interface{}
		var err error

		api := datadogV2.NewAPIManagementApi(client.NewAPIClient())
		res, _, err = api.GetOpenAPI(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to get-open-api")

		cmd.Println(cmdutil.FormatJSON(res, "api_management"))
	},
}

func init() {

	Cmd.AddCommand(GetOpenAPICmd)
}
