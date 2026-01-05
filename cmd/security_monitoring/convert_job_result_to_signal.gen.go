package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var ConvertJobResultToSignalCmd = &cobra.Command{
	Use: "convert-job-result-to-signal [payload]",

	Short: "Convert a job result to a signal",
	Long: `Convert a job result to a signal
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#convert-job-result-to-signal`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.ConvertJobResultsToSignalsRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err = api.ConvertJobResultToSignal(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to convert-job-result-to-signal")

	},
}

func init() {
	Cmd.AddCommand(ConvertJobResultToSignalCmd)
}
