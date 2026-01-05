package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteCustomFrameworkCmd = &cobra.Command{
	Use: "delete-custom-framework [handle] [version]",

	Short: "Delete a custom framework",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.DeleteCustomFramework(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to delete-custom-framework")

		cmd.Println(cmdutil.FormatJSON(res, "custom_framework"))
	},
}

func init() {
	Cmd.AddCommand(DeleteCustomFrameworkCmd)
}
