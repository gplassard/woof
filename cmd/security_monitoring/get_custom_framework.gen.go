package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetCustomFrameworkCmd = &cobra.Command{
	Use: "get-custom-framework [handle] [version]",

	Short: "Get a custom framework",
	Long: `Get a custom framework
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#get-custom-framework`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GetCustomFrameworkResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.GetCustomFramework(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to get-custom-framework")

		cmd.Println(cmdutil.FormatJSON(res, "custom_framework"))
	},
}

func init() {

	Cmd.AddCommand(GetCustomFrameworkCmd)
}
