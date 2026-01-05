package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateSecurityFilterCmd = &cobra.Command{
	Use: "create-security-filter",

	Short: "Create a security filter",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.CreateSecurityFilter(client.NewContext(apiKey, appKey, site), datadogV2.SecurityFilterCreateRequest{})
		cmdutil.HandleError(err, "failed to create-security-filter")

		cmd.Println(cmdutil.FormatJSON(res, "security_filters"))
	},
}

func init() {
	Cmd.AddCommand(CreateSecurityFilterCmd)
}
