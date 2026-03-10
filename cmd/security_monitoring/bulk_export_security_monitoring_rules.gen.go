package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var BulkExportSecurityMonitoringRulesCmd = &cobra.Command{
	Use:     "bulk-export-security-monitoring-rules",
	Aliases: []string{"bulk-export-rules"},
	Short:   "Bulk export security monitoring rules",
	Long: `Bulk export security monitoring rules
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#bulk-export-security-monitoring-rules`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res interface{}
		var err error

		var body datadogV2.SecurityMonitoringRuleBulkExportPayload
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.BulkExportSecurityMonitoringRules(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to bulk-export-security-monitoring-rules")

		fmt.Println(cmdutil.FormatJSON(res, "security_monitoring"))
	},
}

func init() {

	BulkExportSecurityMonitoringRulesCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	BulkExportSecurityMonitoringRulesCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(BulkExportSecurityMonitoringRulesCmd)
}
