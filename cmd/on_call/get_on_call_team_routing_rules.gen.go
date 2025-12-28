package on_call

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetOnCallTeamRoutingRulesCmd = &cobra.Command{
	Use:   "get-on-call-team-routing-rules [team_id]",
	Short: "Get On-Call team routing rules",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.GetOnCallTeamRoutingRules(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get-on-call-team-routing-rules: %v", err)
		}

		cmdutil.PrintJSON(res, "team_routing_rules")
	},
}

func init() {
	Cmd.AddCommand(GetOnCallTeamRoutingRulesCmd)
}
