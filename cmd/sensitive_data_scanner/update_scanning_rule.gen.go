package sensitive_data_scanner

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateScanningRuleCmd = &cobra.Command{
	Use: "update-scanning-rule [rule_id]",

	Short: "Update Scanning Rule",
	Long: `Update Scanning Rule
Documentation: https://docs.datadoghq.com/api/latest/sensitive-data-scanner/#update-scanning-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SensitiveDataScannerRuleUpdateResponse
		var err error

		var body datadogV2.SensitiveDataScannerRuleUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateScanningRule(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-scanning-rule")

		fmt.Println(cmdutil.FormatJSON(res, "sensitive_data_scanner"))
	},
}

func init() {

	UpdateScanningRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateScanningRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateScanningRuleCmd)
}
