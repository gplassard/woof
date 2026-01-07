package case_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCaseCmd = &cobra.Command{
	Use: "create-case",

	Short: "Create a case",
	Long: `Create a case
Documentation: https://docs.datadoghq.com/api/latest/case-management/#create-case`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CaseResponse
		var err error

		var body datadogV2.CaseCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateCase(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-case")

		fmt.Println(cmdutil.FormatJSON(res, "case"))
	},
}

func init() {

	CreateCaseCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateCaseCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateCaseCmd)
}
