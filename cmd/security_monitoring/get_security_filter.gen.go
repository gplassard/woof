package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSecurityFilterCmd = &cobra.Command{
	Use: "get-security-filter [security_filter_id]",

	Short: "Get a security filter",
	Long: `Get a security filter
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#get-security-filter`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityFilterResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.GetSecurityFilter(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-security-filter")

		cmd.Println(cmdutil.FormatJSON(res, "security_filters"))
	},
}

func init() {

	Cmd.AddCommand(GetSecurityFilterCmd)
}
