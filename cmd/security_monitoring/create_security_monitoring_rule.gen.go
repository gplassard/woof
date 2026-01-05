package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateSecurityMonitoringRuleCmd = &cobra.Command{
	Use:     "create-security-monitoring-rule [payload]",
	Aliases: []string{"create-rule"},
	Short:   "Create a detection rule",
	Long: `Create a detection rule
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#create-security-monitoring-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringRuleResponse
		var err error

		var body datadogV2.SecurityMonitoringRuleCreatePayload
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.CreateSecurityMonitoringRule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-security-monitoring-rule")

		cmd.Println(cmdutil.FormatJSON(res, "security_monitoring"))
	},
}

func init() {
	Cmd.AddCommand(CreateSecurityMonitoringRuleCmd)
}
