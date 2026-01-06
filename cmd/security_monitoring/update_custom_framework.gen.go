package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateCustomFrameworkCmd = &cobra.Command{
	Use: "update-custom-framework [handle] [version]",

	Short: "Update a custom framework",
	Long: `Update a custom framework
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#update-custom-framework`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UpdateCustomFrameworkResponse
		var err error

		var body datadogV2.UpdateCustomFrameworkRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateCustomFramework(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-custom-framework")

		cmd.Println(cmdutil.FormatJSON(res, "custom_framework"))
	},
}

func init() {

	UpdateCustomFrameworkCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateCustomFrameworkCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateCustomFrameworkCmd)
}
