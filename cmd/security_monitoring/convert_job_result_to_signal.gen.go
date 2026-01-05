package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ConvertJobResultToSignalCmd = &cobra.Command{
	Use: "convert-job-result-to-signal",

	Short: "Convert a job result to a signal",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err := api.ConvertJobResultToSignal(client.NewContext(apiKey, appKey, site), datadogV2.ConvertJobResultsToSignalsRequest{})
		cmdutil.HandleError(err, "failed to convert-job-result-to-signal")

	},
}

func init() {
	Cmd.AddCommand(ConvertJobResultToSignalCmd)
}
