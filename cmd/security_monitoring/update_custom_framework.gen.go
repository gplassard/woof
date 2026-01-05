package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateCustomFrameworkCmd = &cobra.Command{
	Use: "update-custom-framework [handle] [version] [payload]",

	Short: "Update a custom framework",
	Long: `Update a custom framework
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#update-custom-framework`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UpdateCustomFrameworkResponse
		var err error

		var body datadogV2.UpdateCustomFrameworkRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.UpdateCustomFramework(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-custom-framework")

		cmd.Println(cmdutil.FormatJSON(res, "custom_framework"))
	},
}

func init() {
	Cmd.AddCommand(UpdateCustomFrameworkCmd)
}
