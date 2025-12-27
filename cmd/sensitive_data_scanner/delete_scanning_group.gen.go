package sensitive_data_scanner

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteScanningGroupCmd = &cobra.Command{
	Use:   "delete_scanning_group [group_id]",
	Short: "Delete Scanning Group",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		res, _, err := api.DeleteScanningGroup(client.NewContext(apiKey, appKey, site), args[0], datadogV2.SensitiveDataScannerGroupDeleteRequest{})
		if err != nil {
			log.Fatalf("failed to delete_scanning_group: %v", err)
		}

		cmdutil.PrintJSON(res, "sensitive_data_scanner")
	},
}

func init() {
	Cmd.AddCommand(DeleteScanningGroupCmd)
}
