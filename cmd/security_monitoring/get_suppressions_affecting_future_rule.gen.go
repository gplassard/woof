package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var GetSuppressionsAffectingFutureRuleCmd = &cobra.Command{
	Use: "get-suppressions-affecting-future-rule [payload]",

	Short: "Get suppressions affecting future rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.SecurityMonitoringRuleCreatePayload
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.GetSuppressionsAffectingFutureRule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to get-suppressions-affecting-future-rule")

		cmd.Println(cmdutil.FormatJSON(res, "suppressions"))
	},
}

func init() {
	Cmd.AddCommand(GetSuppressionsAffectingFutureRuleCmd)
}
