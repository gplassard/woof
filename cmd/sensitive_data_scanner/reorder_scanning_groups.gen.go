package sensitive_data_scanner

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var ReorderScanningGroupsCmd = &cobra.Command{
	Use: "reorder-scanning-groups [payload]",

	Short: "Reorder Groups",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SensitiveDataScannerReorderGroupsResponse
		var err error

		var body datadogV2.SensitiveDataScannerConfigRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		res, _, err = api.ReorderScanningGroups(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to reorder-scanning-groups")

		cmd.Println(cmdutil.FormatJSON(res, "sensitive_data_scanner"))
	},
}

func init() {
	Cmd.AddCommand(ReorderScanningGroupsCmd)
}
