package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListVulnerabilitiesCmd = &cobra.Command{
	Use:   "list-vulnerabilities",
	Short: "List vulnerabilities",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListVulnerabilities(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-vulnerabilities: %v", err)
		}

		cmdutil.PrintJSON(res, "vulnerabilities")
	},
}

func init() {
	Cmd.AddCommand(ListVulnerabilitiesCmd)
}
