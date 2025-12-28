package csm_threats

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DownloadCloudWorkloadPolicyFileCmd = &cobra.Command{
	Use:   "download-cloud-workload-policy-file",
	Short: "Download the Workload Protection policy (US1-FED)",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.DownloadCloudWorkloadPolicyFile(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to download-cloud-workload-policy-file: %v", err)
		}

		cmdutil.PrintJSON(res, "csm_threats")
	},
}

func init() {
	Cmd.AddCommand(DownloadCloudWorkloadPolicyFileCmd)
}
