package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCustomFrameworkCmd = &cobra.Command{
	Use: "create-custom-framework",

	Short: "Create a custom framework",
	Long: `Create a custom framework
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#create-custom-framework`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CreateCustomFrameworkResponse
		var err error

		var body datadogV2.CreateCustomFrameworkRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateCustomFramework(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-custom-framework")

		cmd.Println(cmdutil.FormatJSON(res, "custom_framework"))
	},
}

func init() {

	CreateCustomFrameworkCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateCustomFrameworkCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateCustomFrameworkCmd)
}
