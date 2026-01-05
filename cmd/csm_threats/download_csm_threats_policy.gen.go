package csm_threats

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DownloadCSMThreatsPolicyCmd = &cobra.Command{
	Use:     "download-csm-threats-policy",
	Aliases: []string{"download-policy"},
	Short:   "Download the Workload Protection policy",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.DownloadCSMThreatsPolicy(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to download-csm-threats-policy")

		cmd.Println(cmdutil.FormatJSON(res, "csm_threats"))
	},
}

func init() {
	Cmd.AddCommand(DownloadCSMThreatsPolicyCmd)
}
