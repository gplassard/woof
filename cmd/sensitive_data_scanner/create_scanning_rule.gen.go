package sensitive_data_scanner

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateScanningRuleCmd = &cobra.Command{
	Use:   "create_scanning_rule",
	Short: "Create Scanning Rule",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		res, _, err := api.CreateScanningRule(client.NewContext(apiKey, appKey, site), datadogV2.SensitiveDataScannerRuleCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create_scanning_rule: %v", err)
		}

		cmdutil.PrintJSON(res, "sensitive_data_scanner")
	},
}

func init() {
	Cmd.AddCommand(CreateScanningRuleCmd)
}
