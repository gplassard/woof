package case_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetCaseCmd = &cobra.Command{
	Use: "get-case [case_id]",

	Short: "Get the details of a case",
	Long: `Get the details of a case
Documentation: https://docs.datadoghq.com/api/latest/case-management/#get-case`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CaseResponse
		var err error

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err = api.GetCase(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-case")

		cmd.Println(cmdutil.FormatJSON(res, "case"))
	},
}

func init() {
	Cmd.AddCommand(GetCaseCmd)
}
