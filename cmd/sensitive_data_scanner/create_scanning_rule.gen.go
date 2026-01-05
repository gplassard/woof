package sensitive_data_scanner

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateScanningRuleCmd = &cobra.Command{
	Use: "create-scanning-rule [payload]",

	Short: "Create Scanning Rule",
	Long: `Create Scanning Rule
Documentation: https://docs.datadoghq.com/api/latest/sensitive-data-scanner/#create-scanning-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SensitiveDataScannerCreateRuleResponse
		var err error

		var body datadogV2.SensitiveDataScannerRuleCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		res, _, err = api.CreateScanningRule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-scanning-rule")

		cmd.Println(cmdutil.FormatJSON(res, "sensitive_data_scanner_rule"))
	},
}

func init() {
	Cmd.AddCommand(CreateScanningRuleCmd)
}
