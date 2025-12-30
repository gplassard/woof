package security_monitoring

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSuppressionsAffectingFutureRuleCmd = &cobra.Command{
	Use: "get-suppressions-affecting-future-rule",

	Short: "Get suppressions affecting future rule",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.GetSuppressionsAffectingFutureRule(client.NewContext(apiKey, appKey, site), datadogV2.SecurityMonitoringRuleCreatePayload{})
		cmdutil.HandleError(err, "failed to get-suppressions-affecting-future-rule")

		cmd.Println(cmdutil.FormatJSON(res, "suppressions"))
	},
}

func init() {
	Cmd.AddCommand(GetSuppressionsAffectingFutureRuleCmd)
}
