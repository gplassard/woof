package sensitive_data_scanner

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateScanningGroupCmd = &cobra.Command{
	Use: "create-scanning-group",

	Short: "Create Scanning Group",
	Long: `Create Scanning Group
Documentation: https://docs.datadoghq.com/api/latest/sensitive-data-scanner/#create-scanning-group`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SensitiveDataScannerCreateGroupResponse
		var err error

		var body datadogV2.SensitiveDataScannerGroupCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateScanningGroup(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-scanning-group")

		fmt.Println(cmdutil.FormatJSON(res, "sensitive_data_scanner_group"))
	},
}

func init() {

	CreateScanningGroupCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateScanningGroupCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateScanningGroupCmd)
}
