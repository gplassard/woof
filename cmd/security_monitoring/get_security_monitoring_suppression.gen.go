package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSecurityMonitoringSuppressionCmd = &cobra.Command{
	Use:     "get-security-monitoring-suppression [suppression_id]",
	Aliases: []string{"get-suppression"},
	Short:   "Get a suppression rule",
	Long: `Get a suppression rule
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#get-security-monitoring-suppression`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringSuppressionResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetSecurityMonitoringSuppression(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-security-monitoring-suppression")

		cmd.Println(cmdutil.FormatJSON(res, "suppressions"))
	},
}

func init() {

	Cmd.AddCommand(GetSecurityMonitoringSuppressionCmd)
}
