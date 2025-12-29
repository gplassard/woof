package sensitive_data_scanner

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListScanningGroupsCmd = &cobra.Command{
	Use:   "list-scanning-groups",
	
	Short: "List Scanning Groups",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		res, _, err := api.ListScanningGroups(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-scanning-groups")

		cmd.Println(cmdutil.FormatJSON(res, "sensitive_data_scanner_configuration"))
	},
}

func init() {
	Cmd.AddCommand(ListScanningGroupsCmd)
}
