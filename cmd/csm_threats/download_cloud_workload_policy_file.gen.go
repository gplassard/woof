package csm_threats

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DownloadCloudWorkloadPolicyFileCmd = &cobra.Command{
	Use: "download-cloud-workload-policy-file",

	Short: "Download the Workload Protection policy (US1-FED)",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.DownloadCloudWorkloadPolicyFile(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to download-cloud-workload-policy-file")

		cmd.Println(cmdutil.FormatJSON(res, "csm_threats"))
	},
}

func init() {
	Cmd.AddCommand(DownloadCloudWorkloadPolicyFileCmd)
}
