package security_monitoring

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ValidateSecurityMonitoringSuppressionCmd = &cobra.Command{
	Use:   "validate-security-monitoring-suppression",
	Aliases: []string{ "validate-suppression", },
	Short: "Validate a suppression rule",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err := api.ValidateSecurityMonitoringSuppression(client.NewContext(apiKey, appKey, site), datadogV2.SecurityMonitoringSuppressionCreateRequest{})
		cmdutil.HandleError(err, "failed to validate-security-monitoring-suppression")

		
	},
}

func init() {
	Cmd.AddCommand(ValidateSecurityMonitoringSuppressionCmd)
}
