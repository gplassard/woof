package sensitive_data_scanner

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateScanningRuleCmd = &cobra.Command{
	Use: "create-scanning-rule",

	Short: "Create Scanning Rule",
	Long: `Create Scanning Rule
Documentation: https://docs.datadoghq.com/api/latest/sensitive-data-scanner/#create-scanning-rule`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SensitiveDataScannerCreateRuleResponse
		var err error

		var body datadogV2.SensitiveDataScannerRuleCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateScanningRule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-scanning-rule")

		fmt.Println(cmdutil.FormatJSON(res, "scanning_rule"))
	},
}

func init() {

	CreateScanningRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateScanningRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateScanningRuleCmd)
}
