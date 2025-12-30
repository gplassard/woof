package sensitive_data_scanner

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateScanningGroupCmd = &cobra.Command{
	Use: "create-scanning-group",

	Short: "Create Scanning Group",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		res, _, err := api.CreateScanningGroup(client.NewContext(apiKey, appKey, site), datadogV2.SensitiveDataScannerGroupCreateRequest{})
		cmdutil.HandleError(err, "failed to create-scanning-group")

		cmd.Println(cmdutil.FormatJSON(res, "sensitive_data_scanner_group"))
	},
}

func init() {
	Cmd.AddCommand(CreateScanningGroupCmd)
}
