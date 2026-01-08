package csm_threats

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"io"
)

var DownloadCloudWorkloadPolicyFileCmd = &cobra.Command{
	Use: "download-cloud-workload-policy-file",

	Short: "Download the Workload Protection policy (US1-FED)",
	Long: `Download the Workload Protection policy (US1-FED)
Documentation: https://docs.datadoghq.com/api/latest/csm-threats/#download-cloud-workload-policy-file`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res io.Reader
		var err error

		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.DownloadCloudWorkloadPolicyFile(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to download-cloud-workload-policy-file")

		fmt.Println(cmdutil.FormatJSON(res, "csm_threats"))
	},
}

func init() {

	Cmd.AddCommand(DownloadCloudWorkloadPolicyFileCmd)
}
