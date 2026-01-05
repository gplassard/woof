package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DetachCaseCmd = &cobra.Command{
	Use: "detach-case",

	Short: "Detach security findings from their case",
	Long: `Detach security findings from their case
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#detach-case`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.DetachCaseRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err = api.DetachCase(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to detach-case")

	},
}

func init() {

	DetachCaseCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	DetachCaseCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(DetachCaseCmd)
}
