package sensitive_data_scanner

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListScanningGroupsCmd = &cobra.Command{
	Use:   "list_scanning_groups",
	Short: "List Scanning Groups",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		res, _, err := api.ListScanningGroups(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list_scanning_groups: %v", err)
		}

		cmdutil.PrintJSON(res, "sensitive_data_scanner")
	},
}

func init() {
	Cmd.AddCommand(ListScanningGroupsCmd)
}
