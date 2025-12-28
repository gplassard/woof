package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetSecretsRulesCmd = &cobra.Command{
	Use:   "get-secrets-rules",
	Short: "Returns a list of Secrets rules",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.GetSecretsRules(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to get-secrets-rules: %v", err)
		}

		cmdutil.PrintJSON(res, "secret_rule")
	},
}

func init() {
	Cmd.AddCommand(GetSecretsRulesCmd)
}
