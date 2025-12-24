package security_monitoring

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var MuteFindingsCmd = &cobra.Command{
	Use:   "mutefindings",
	Short: "Mute or unmute a batch of findings",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.MuteFindings(client.NewContext(apiKey, appKey, site), datadogV2.BulkMuteFindingsRequest{})
		if err != nil {
			log.Fatalf("failed to mutefindings: %v", err)
		}

		cmdutil.PrintJSON(res, "security_monitoring")
	},
}

func init() {
	Cmd.AddCommand(MuteFindingsCmd)
}
