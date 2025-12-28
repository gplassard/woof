package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetSecurityFilterCmd = &cobra.Command{
	Use:   "get-security-filter [security_filter_id]",
	Short: "Get a security filter",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.GetSecurityFilter(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get-security-filter: %v", err)
		}

		cmdutil.PrintJSON(res, "security_filters")
	},
}

func init() {
	Cmd.AddCommand(GetSecurityFilterCmd)
}
