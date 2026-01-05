package sensitive_data_scanner

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteScanningRuleCmd = &cobra.Command{
	Use: "delete-scanning-rule [rule_id]",

	Short: "Delete Scanning Rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		res, _, err := api.DeleteScanningRule(client.NewContext(apiKey, appKey, site), args[0], datadogV2.SensitiveDataScannerRuleDeleteRequest{})
		cmdutil.HandleError(err, "failed to delete-scanning-rule")

		cmd.Println(cmdutil.FormatJSON(res, "sensitive_data_scanner"))
	},
}

func init() {
	Cmd.AddCommand(DeleteScanningRuleCmd)
}
