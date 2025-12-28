package security_monitoring

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ConvertSecurityMonitoringRuleFromJSONToTerraformCmd = &cobra.Command{
	Use:   "convert-security-monitoring-rule-from-json-to-terraform",
	Aliases: []string{ "convert-rule-from-json-to-terraform", },
	Short: "Convert a rule from JSON to Terraform",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ConvertSecurityMonitoringRuleFromJSONToTerraform(client.NewContext(apiKey, appKey, site), datadogV2.SecurityMonitoringRuleConvertPayload{})
		cmdutil.HandleError(err, "failed to convert-security-monitoring-rule-from-json-to-terraform")

		cmdutil.PrintJSON(res, "security_monitoring")
	},
}

func init() {
	Cmd.AddCommand(ConvertSecurityMonitoringRuleFromJSONToTerraformCmd)
}
