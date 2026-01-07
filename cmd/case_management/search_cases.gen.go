package case_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SearchCasesCmd = &cobra.Command{
	Use: "search-cases",

	Short: "Search cases",
	Long: `Search cases
Documentation: https://docs.datadoghq.com/api/latest/case-management/#search-cases`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CasesResponse
		var err error

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.SearchCases(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to search-cases")

		fmt.Println(cmdutil.FormatJSON(res, "case"))
	},
}

func init() {

	Cmd.AddCommand(SearchCasesCmd)
}
