package security_monitoring

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var AttachCaseCmd = &cobra.Command{
	Use: "attach-case [case_id]",

	Short: "Attach security findings to a case",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.AttachCase(client.NewContext(apiKey, appKey, site), args[0], datadogV2.AttachCaseRequest{})
		cmdutil.HandleError(err, "failed to attach-case")

		cmd.Println(cmdutil.FormatJSON(res, "cases"))
	},
}

func init() {
	Cmd.AddCommand(AttachCaseCmd)
}
