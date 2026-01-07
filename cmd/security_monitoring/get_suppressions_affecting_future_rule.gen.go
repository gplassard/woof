package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSuppressionsAffectingFutureRuleCmd = &cobra.Command{
	Use: "get-suppressions-affecting-future-rule",

	Short: "Get suppressions affecting future rule",
	Long: `Get suppressions affecting future rule
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#get-suppressions-affecting-future-rule`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringSuppressionsResponse
		var err error

		var body datadogV2.SecurityMonitoringRuleCreatePayload
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetSuppressionsAffectingFutureRule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to get-suppressions-affecting-future-rule")

		fmt.Println(cmdutil.FormatJSON(res, "suppressions"))
	},
}

func init() {

	GetSuppressionsAffectingFutureRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	GetSuppressionsAffectingFutureRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(GetSuppressionsAffectingFutureRuleCmd)
}
