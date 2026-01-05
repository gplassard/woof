package security_monitoring

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var MuteFindingsCmd = &cobra.Command{
	Use: "mute-findings",

	Short: "Mute or unmute a batch of findings",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.MuteFindings(client.NewContext(apiKey, appKey, site), datadogV2.BulkMuteFindingsRequest{})
		cmdutil.HandleError(err, "failed to mute-findings")

		cmd.Println(cmdutil.FormatJSON(res, "finding"))
	},
}

func init() {
	Cmd.AddCommand(MuteFindingsCmd)
}
