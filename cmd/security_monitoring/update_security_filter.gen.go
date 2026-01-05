package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateSecurityFilterCmd = &cobra.Command{
	Use: "update-security-filter [security_filter_id]",

	Short: "Update a security filter",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.UpdateSecurityFilter(client.NewContext(apiKey, appKey, site), args[0], datadogV2.SecurityFilterUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-security-filter")

		cmd.Println(cmdutil.FormatJSON(res, "security_filters"))
	},
}

func init() {
	Cmd.AddCommand(UpdateSecurityFilterCmd)
}
