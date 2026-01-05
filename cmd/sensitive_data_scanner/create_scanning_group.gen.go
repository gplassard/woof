package sensitive_data_scanner

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateScanningGroupCmd = &cobra.Command{
	Use: "create-scanning-group [payload]",

	Short: "Create Scanning Group",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.SensitiveDataScannerGroupCreateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		res, _, err := api.CreateScanningGroup(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-scanning-group")

		cmd.Println(cmdutil.FormatJSON(res, "sensitive_data_scanner_group"))
	},
}

func init() {
	Cmd.AddCommand(CreateScanningGroupCmd)
}
