package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListSecurityFiltersCmd = &cobra.Command{
	Use:   "list-security-filters",
	
	Short: "Get all security filters",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListSecurityFilters(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-security-filters: %v", err)
		}

		cmdutil.PrintJSON(res, "security_filters")
	},
}

func init() {
	Cmd.AddCommand(ListSecurityFiltersCmd)
}
