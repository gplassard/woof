package security_monitoring

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ConvertExistingSecurityMonitoringRuleCmd = &cobra.Command{
	Use:   "convert-existing-security-monitoring-rule [rule_id]",
	Aliases: []string{ "convert-existing-rule", },
	Short: "Convert an existing rule from JSON to Terraform",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ConvertExistingSecurityMonitoringRule(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to convert-existing-security-monitoring-rule")

		cmd.Println(cmdutil.FormatJSON(res, "security_monitoring"))
	},
}

func init() {
	Cmd.AddCommand(ConvertExistingSecurityMonitoringRuleCmd)
}
