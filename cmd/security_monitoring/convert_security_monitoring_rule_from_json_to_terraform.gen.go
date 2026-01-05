package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var ConvertSecurityMonitoringRuleFromJSONToTerraformCmd = &cobra.Command{
	Use:     "convert-security-monitoring-rule-from-json-to-terraform [payload]",
	Aliases: []string{"convert-rule-from-json-to-terraform"},
	Short:   "Convert a rule from JSON to Terraform",
	Long: `Convert a rule from JSON to Terraform
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#convert-security-monitoring-rule-from-json-to-terraform`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringRuleConvertResponse
		var err error

		var body datadogV2.SecurityMonitoringRuleConvertPayload
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.ConvertSecurityMonitoringRuleFromJSONToTerraform(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to convert-security-monitoring-rule-from-json-to-terraform")

		cmd.Println(cmdutil.FormatJSON(res, "security_monitoring"))
	},
}

func init() {
	Cmd.AddCommand(ConvertSecurityMonitoringRuleFromJSONToTerraformCmd)
}
