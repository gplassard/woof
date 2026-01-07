package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateSecurityMonitoringRuleCmd = &cobra.Command{
	Use:     "create-security-monitoring-rule",
	Aliases: []string{"create-rule"},
	Short:   "Create a detection rule",
	Long: `Create a detection rule
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#create-security-monitoring-rule`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringRuleResponse
		var err error

		var body datadogV2.SecurityMonitoringRuleCreatePayload
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateSecurityMonitoringRule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-security-monitoring-rule")

		fmt.Println(cmdutil.FormatJSON(res, "security_monitoring"))
	},
}

func init() {

	CreateSecurityMonitoringRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateSecurityMonitoringRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateSecurityMonitoringRuleCmd)
}
