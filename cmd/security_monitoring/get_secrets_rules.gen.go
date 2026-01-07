package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSecretsRulesCmd = &cobra.Command{
	Use: "get-secrets-rules",

	Short: "Returns a list of Secrets rules",
	Long: `Returns a list of Secrets rules
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#get-secrets-rules`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecretRuleArray
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetSecretsRules(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-secrets-rules")

		fmt.Println(cmdutil.FormatJSON(res, "secret_rule"))
	},
}

func init() {

	Cmd.AddCommand(GetSecretsRulesCmd)
}
