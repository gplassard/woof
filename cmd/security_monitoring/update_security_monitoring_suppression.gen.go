package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateSecurityMonitoringSuppressionCmd = &cobra.Command{
	Use:     "update-security-monitoring-suppression [suppression_id]",
	Aliases: []string{"update-suppression"},
	Short:   "Update a suppression rule",
	Long: `Update a suppression rule
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#update-security-monitoring-suppression`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringSuppressionResponse
		var err error

		var body datadogV2.SecurityMonitoringSuppressionUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateSecurityMonitoringSuppression(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-security-monitoring-suppression")

		cmd.Println(cmdutil.FormatJSON(res, "suppressions"))
	},
}

func init() {

	UpdateSecurityMonitoringSuppressionCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateSecurityMonitoringSuppressionCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateSecurityMonitoringSuppressionCmd)
}
