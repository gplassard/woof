package security_monitoring

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListVulnerableAssetsCmd = &cobra.Command{
	Use:   "listvulnerableassets",
	Short: "List vulnerable assets",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListVulnerableAssets(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listvulnerableassets: %v", err)
		}

		cmdutil.PrintJSON(res, "security_monitoring")
	},
}

func init() {
	Cmd.AddCommand(ListVulnerableAssetsCmd)
}
