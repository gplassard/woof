package sensitive_data_scanner

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateScanningRuleCmd = &cobra.Command{
	Use:   "update-scanning-rule [rule_id]",
	
	Short: "Update Scanning Rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		res, _, err := api.UpdateScanningRule(client.NewContext(apiKey, appKey, site), args[0], datadogV2.SensitiveDataScannerRuleUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-scanning-rule")

		cmdutil.PrintJSON(res, "sensitive_data_scanner")
	},
}

func init() {
	Cmd.AddCommand(UpdateScanningRuleCmd)
}
