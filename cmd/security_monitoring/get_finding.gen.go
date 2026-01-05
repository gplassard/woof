package security_monitoring

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetFindingCmd = &cobra.Command{
	Use: "get-finding [finding_id]",

	Short: "Get a finding",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.GetFinding(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-finding")

		cmd.Println(cmdutil.FormatJSON(res, "detailed_finding"))
	},
}

func init() {
	Cmd.AddCommand(GetFindingCmd)
}
