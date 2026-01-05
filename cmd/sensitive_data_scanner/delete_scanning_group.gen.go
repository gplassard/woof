package sensitive_data_scanner

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var DeleteScanningGroupCmd = &cobra.Command{
	Use: "delete-scanning-group [group_id] [payload]",

	Short: "Delete Scanning Group",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SensitiveDataScannerGroupDeleteResponse
		var err error

		var body datadogV2.SensitiveDataScannerGroupDeleteRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		res, _, err = api.DeleteScanningGroup(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to delete-scanning-group")

		cmd.Println(cmdutil.FormatJSON(res, "sensitive_data_scanner"))
	},
}

func init() {
	Cmd.AddCommand(DeleteScanningGroupCmd)
}
