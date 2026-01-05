package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ValidateSecurityMonitoringSuppressionCmd = &cobra.Command{
	Use:     "validate-security-monitoring-suppression",
	Aliases: []string{"validate-suppression"},
	Short:   "Validate a suppression rule",
	Long: `Validate a suppression rule
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#validate-security-monitoring-suppression`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.SecurityMonitoringSuppressionCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err = api.ValidateSecurityMonitoringSuppression(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to validate-security-monitoring-suppression")

	},
}

func init() {

	ValidateSecurityMonitoringSuppressionCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	ValidateSecurityMonitoringSuppressionCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(ValidateSecurityMonitoringSuppressionCmd)
}
