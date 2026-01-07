package sensitive_data_scanner

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ReorderScanningGroupsCmd = &cobra.Command{
	Use: "reorder-scanning-groups",

	Short: "Reorder Groups",
	Long: `Reorder Groups
Documentation: https://docs.datadoghq.com/api/latest/sensitive-data-scanner/#reorder-scanning-groups`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SensitiveDataScannerReorderGroupsResponse
		var err error

		var body datadogV2.SensitiveDataScannerConfigRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ReorderScanningGroups(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to reorder-scanning-groups")

		fmt.Println(cmdutil.FormatJSON(res, "sensitive_data_scanner"))
	},
}

func init() {

	ReorderScanningGroupsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	ReorderScanningGroupsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(ReorderScanningGroupsCmd)
}
