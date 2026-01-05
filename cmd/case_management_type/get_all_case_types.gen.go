package case_management_type

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetAllCaseTypesCmd = &cobra.Command{
	Use: "get-all-case-types",

	Short: "Get all case types",
	Long: `Get all case types
Documentation: https://docs.datadoghq.com/api/latest/case-management-type/#get-all-case-types`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CaseTypesResponse
		var err error

		api := datadogV2.NewCaseManagementTypeApi(client.NewAPIClient())
		res, _, err = api.GetAllCaseTypes(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-all-case-types")

		cmd.Println(cmdutil.FormatJSON(res, "case_type"))
	},
}

func init() {
	Cmd.AddCommand(GetAllCaseTypesCmd)
}
