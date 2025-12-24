package security_monitoring

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListVulnerabilitiesCmd = &cobra.Command{
	Use:   "listvulnerabilities",
	Short: "List vulnerabilities",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListVulnerabilities(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listvulnerabilities: %v", err)
		}

		cmdutil.PrintJSON(res, "security_monitoring")
	},
}

func init() {
	Cmd.AddCommand(ListVulnerabilitiesCmd)
}
