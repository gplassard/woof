package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ConvertSecurityMonitoringRuleFromJSONToTerraformCmd = &cobra.Command{
	Use:     "convert-security-monitoring-rule-from-json-to-terraform",
	Aliases: []string{"convert-rule-from-json-to-terraform"},
	Short:   "Convert a rule from JSON to Terraform",
	Long: `Convert a rule from JSON to Terraform
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#convert-security-monitoring-rule-from-json-to-terraform`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringRuleConvertResponse
		var err error

		var body datadogV2.SecurityMonitoringRuleConvertPayload
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ConvertSecurityMonitoringRuleFromJSONToTerraform(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to convert-security-monitoring-rule-from-json-to-terraform")

		fmt.Println(cmdutil.FormatJSON(res, "security_monitoring"))
	},
}

func init() {

	ConvertSecurityMonitoringRuleFromJSONToTerraformCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	ConvertSecurityMonitoringRuleFromJSONToTerraformCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(ConvertSecurityMonitoringRuleFromJSONToTerraformCmd)
}
