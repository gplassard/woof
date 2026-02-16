package api_management

import (
	"fmt"
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

		optionalParams := datadogV2.NewListAPIsOptionalParameters()

		if cmd.Flags().Changed("query") {
			val, _ := cmd.Flags().GetString("query")
			optionalParams.WithQuery(val)
		}

		if cmd.Flags().Changed("page-limit") {
			val, _ := cmd.Flags().GetInt64("page-limit")
			optionalParams.WithPageLimit(val)
		}

		if cmd.Flags().Changed("page-offset") {
			val, _ := cmd.Flags().GetInt64("page-offset")
			optionalParams.WithPageOffset(val)
		}

		api := datadogV2.NewAPIManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListAPIs(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-apis")

		fmt.Println(cmdutil.FormatJSON(res, "api_management"))
	},
}

func init() {

	ListAPIsCmd.Flags().String("query", "", "Filter APIs by name")

	ListAPIsCmd.Flags().Int64("page-limit", 0, "Number of items per page.")

	ListAPIsCmd.Flags().Int64("page-offset", 0, "Offset for pagination.")

	Cmd.AddCommand(ListAPIsCmd)
}
