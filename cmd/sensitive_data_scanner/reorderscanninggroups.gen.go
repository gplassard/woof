package sensitive_data_scanner

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ReorderScanningGroupsCmd = &cobra.Command{
	Use:   "reorderscanninggroups",
	Short: "Reorder Groups",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		res, _, err := api.ReorderScanningGroups(client.NewContext(apiKey, appKey, site), datadogV2.SensitiveDataScannerConfigRequest{})
		if err != nil {
			log.Fatalf("failed to reorderscanninggroups: %v", err)
		}

		cmdutil.PrintJSON(res, "sensitive_data_scanner")
	},
}

func init() {
	Cmd.AddCommand(ReorderScanningGroupsCmd)
}
