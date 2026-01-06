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
	Long: `Update a security filter
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#update-security-filter`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityFilterResponse
		var err error

		var body datadogV2.SecurityFilterUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateSecurityFilter(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-security-filter")

		cmd.Println(cmdutil.FormatJSON(res, "security_filters"))
	},
}

func init() {

	UpdateSecurityFilterCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateSecurityFilterCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateSecurityFilterCmd)
}
