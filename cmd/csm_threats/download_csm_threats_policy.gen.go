package csm_threats

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DownloadCSMThreatsPolicyCmd = &cobra.Command{
	Use:   "download-csm-threats-policy",
	Aliases: []string{ "download-policy", },
	Short: "Download the Workload Protection policy",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.DownloadCSMThreatsPolicy(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to download-csm-threats-policy")

		cmdutil.PrintJSON(res, "csm_threats")
	},
}

func init() {
	Cmd.AddCommand(DownloadCSMThreatsPolicyCmd)
}
