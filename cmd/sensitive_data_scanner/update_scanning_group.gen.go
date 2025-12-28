package sensitive_data_scanner

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateScanningGroupCmd = &cobra.Command{
	Use:   "update-scanning-group [group_id]",
	Short: "Update Scanning Group",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		res, _, err := api.UpdateScanningGroup(client.NewContext(apiKey, appKey, site), args[0], datadogV2.SensitiveDataScannerGroupUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to update-scanning-group: %v", err)
		}

		cmdutil.PrintJSON(res, "sensitive_data_scanner")
	},
}

func init() {
	Cmd.AddCommand(UpdateScanningGroupCmd)
}
