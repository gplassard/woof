package security_monitoring

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ConvertJobResultToSignalCmd = &cobra.Command{
	Use:   "convertjobresulttosignal",
	Short: "Convert a job result to a signal",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err := api.ConvertJobResultToSignal(client.NewContext(apiKey, appKey, site), datadogV2.ConvertJobResultsToSignalsRequest{})
		if err != nil {
			log.Fatalf("failed to convertjobresulttosignal: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(ConvertJobResultToSignalCmd)
}
