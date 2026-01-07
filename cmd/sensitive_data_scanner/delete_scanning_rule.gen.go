package sensitive_data_scanner

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteScanningRuleCmd = &cobra.Command{
	Use: "delete-scanning-rule [rule_id]",

	Short: "Delete Scanning Rule",
	Long: `Delete Scanning Rule
Documentation: https://docs.datadoghq.com/api/latest/sensitive-data-scanner/#delete-scanning-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SensitiveDataScannerRuleDeleteResponse
		var err error

		var body datadogV2.SensitiveDataScannerRuleDeleteRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.DeleteScanningRule(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to delete-scanning-rule")

		fmt.Println(cmdutil.FormatJSON(res, "sensitive_data_scanner"))
	},
}

func init() {

	DeleteScanningRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	DeleteScanningRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(DeleteScanningRuleCmd)
}
