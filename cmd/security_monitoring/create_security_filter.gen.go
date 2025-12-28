package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateSecurityFilterCmd = &cobra.Command{
	Use:   "create_security_filter",
	Short: "Create a security filter",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.CreateSecurityFilter(client.NewContext(apiKey, appKey, site), datadogV2.SecurityFilterCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create_security_filter: %v", err)
		}

		cmdutil.PrintJSON(res, "security_filters")
	},
}

func init() {
	Cmd.AddCommand(CreateSecurityFilterCmd)
}
