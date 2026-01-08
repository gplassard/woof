package c_s_m_threats

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"io"
)

var DownloadCSMThreatsPolicyCmd = &cobra.Command{
	Use: "download-csm-threats-policy",

	Short: "Download the Workload Protection policy",
	Long: `Download the Workload Protection policy
Documentation: https://docs.datadoghq.com/api/latest/c-s-m-threats/#download-csm-threats-policy`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res io.Reader
		var err error

		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.DownloadCSMThreatsPolicy(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to download-csm-threats-policy")

		fmt.Println(cmdutil.FormatJSON(res, "c_s_m_threats"))
	},
}

func init() {

	Cmd.AddCommand(DownloadCSMThreatsPolicyCmd)
}
