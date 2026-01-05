package sensitive_data_scanner

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateScanningRuleCmd = &cobra.Command{
	Use: "create-scanning-rule",

	Short: "Create Scanning Rule",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		res, _, err := api.CreateScanningRule(client.NewContext(apiKey, appKey, site), datadogV2.SensitiveDataScannerRuleCreateRequest{})
		cmdutil.HandleError(err, "failed to create-scanning-rule")

		cmd.Println(cmdutil.FormatJSON(res, "sensitive_data_scanner_rule"))
	},
}

func init() {
	Cmd.AddCommand(CreateScanningRuleCmd)
}
