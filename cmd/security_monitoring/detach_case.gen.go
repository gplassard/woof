package security_monitoring

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DetachCaseCmd = &cobra.Command{
	Use: "detach-case",

	Short: "Detach security findings from their case",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err := api.DetachCase(client.NewContext(apiKey, appKey, site), datadogV2.DetachCaseRequest{})
		cmdutil.HandleError(err, "failed to detach-case")

	},
}

func init() {
	Cmd.AddCommand(DetachCaseCmd)
}
