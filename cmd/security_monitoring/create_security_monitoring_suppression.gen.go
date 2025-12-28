package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateSecurityMonitoringSuppressionCmd = &cobra.Command{
	Use:   "create-security-monitoring-suppression",
	Short: "Create a suppression rule",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.CreateSecurityMonitoringSuppression(client.NewContext(apiKey, appKey, site), datadogV2.SecurityMonitoringSuppressionCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create-security-monitoring-suppression: %v", err)
		}

		cmdutil.PrintJSON(res, "suppressions")
	},
}

func init() {
	Cmd.AddCommand(CreateSecurityMonitoringSuppressionCmd)
}
