package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var ValidateSecurityMonitoringSuppressionCmd = &cobra.Command{
	Use:     "validate-security-monitoring-suppression [payload]",
	Aliases: []string{"validate-suppression"},
	Short:   "Validate a suppression rule",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.SecurityMonitoringSuppressionCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err = api.ValidateSecurityMonitoringSuppression(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to validate-security-monitoring-suppression")

	},
}

func init() {
	Cmd.AddCommand(ValidateSecurityMonitoringSuppressionCmd)
}
