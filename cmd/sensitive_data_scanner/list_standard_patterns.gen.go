package sensitive_data_scanner

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListStandardPatternsCmd = &cobra.Command{
	Use:   "list-standard-patterns",
	
	Short: "List standard patterns",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSensitiveDataScannerApi(client.NewAPIClient())
		res, _, err := api.ListStandardPatterns(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-standard-patterns")

		cmdutil.PrintJSON(res, "sensitive_data_scanner")
	},
}

func init() {
	Cmd.AddCommand(ListStandardPatternsCmd)
}
