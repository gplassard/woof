package sensitive_data_scanner

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateScanningGroupCmd = &cobra.Command{
	Use: "update-scanning-group [group_id]",

	Short: "Update Scanning Group",
	Long: `Update Scanning Group
Documentation: https://docs.datadoghq.com/api/latest/sensitive-data-scanner/#update-scanning-group`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SensitiveDataScannerGroupUpdateResponse
		var err error

		var body datadogV2.SensitiveDataScannerGroupUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateScanningGroup(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-scanning-group")

		cmd.Println(cmdutil.FormatJSON(res, "sensitive_data_scanner"))
	},
}

func init() {

	UpdateScanningGroupCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateScanningGroupCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateScanningGroupCmd)
}
