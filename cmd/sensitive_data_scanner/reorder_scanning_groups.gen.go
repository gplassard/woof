package sensitive_data_scanner

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ReorderScanningGroupsCmd = &cobra.Command{
	Use: "reorder-scanning-groups",

	Short: "Reorder Groups",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		res, _, err := api.ReorderScanningGroups(client.NewContext(apiKey, appKey, site), datadogV2.SensitiveDataScannerConfigRequest{})
		cmdutil.HandleError(err, "failed to reorder-scanning-groups")

		cmd.Println(cmdutil.FormatJSON(res, "sensitive_data_scanner"))
	},
}

func init() {
	Cmd.AddCommand(ReorderScanningGroupsCmd)
}
