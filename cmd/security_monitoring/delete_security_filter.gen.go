package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteSecurityFilterCmd = &cobra.Command{
	Use: "delete-security-filter [security_filter_id]",

	Short: "Delete a security filter",
	Long: `Delete a security filter
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#delete-security-filter`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err = api.DeleteSecurityFilter(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-security-filter")

	},
}

func init() {

	Cmd.AddCommand(DeleteSecurityFilterCmd)
}
