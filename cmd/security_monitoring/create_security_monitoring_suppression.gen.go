package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateSecurityMonitoringSuppressionCmd = &cobra.Command{
	Use:     "create-security-monitoring-suppression",
	Aliases: []string{"create-suppression"},
	Short:   "Create a suppression rule",
	Long: `Create a suppression rule
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#create-security-monitoring-suppression`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringSuppressionResponse
		var err error

		var body datadogV2.SecurityMonitoringSuppressionCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateSecurityMonitoringSuppression(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-security-monitoring-suppression")

		cmd.Println(cmdutil.FormatJSON(res, "suppressions"))
	},
}

func init() {

	CreateSecurityMonitoringSuppressionCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateSecurityMonitoringSuppressionCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateSecurityMonitoringSuppressionCmd)
}
