package sensitive_data_scanner

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListScanningGroupsCmd = &cobra.Command{
	Use: "list-scanning-groups",

	Short: "List Scanning Groups",
	Long: `List Scanning Groups
Documentation: https://docs.datadoghq.com/api/latest/sensitive-data-scanner/#list-scanning-groups`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SensitiveDataScannerGetConfigResponse
		var err error

		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListScanningGroups(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-scanning-groups")

		fmt.Println(cmdutil.FormatJSON(res, "sensitive_data_scanner_configuration"))
	},
}

func init() {

	Cmd.AddCommand(ListScanningGroupsCmd)
}
