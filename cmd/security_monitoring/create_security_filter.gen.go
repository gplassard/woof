package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateSecurityFilterCmd = &cobra.Command{
	Use: "create-security-filter",

	Short: "Create a security filter",
	Long: `Create a security filter
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#create-security-filter`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityFilterResponse
		var err error

		var body datadogV2.SecurityFilterCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateSecurityFilter(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-security-filter")

		fmt.Println(cmdutil.FormatJSON(res, "security_filters"))
	},
}

func init() {

	CreateSecurityFilterCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateSecurityFilterCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateSecurityFilterCmd)
}
