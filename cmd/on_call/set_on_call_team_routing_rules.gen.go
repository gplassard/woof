package on_call

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var SetOnCallTeamRoutingRulesCmd = &cobra.Command{
	Use:   "set_on_call_team_routing_rules [team_id]",
	Short: "Set On-Call team routing rules",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.SetOnCallTeamRoutingRules(client.NewContext(apiKey, appKey, site), args[0], datadogV2.TeamRoutingRulesRequest{})
		if err != nil {
			log.Fatalf("failed to set_on_call_team_routing_rules: %v", err)
		}

		cmdutil.PrintJSON(res, "on_call")
	},
}

func init() {
	Cmd.AddCommand(SetOnCallTeamRoutingRulesCmd)
}
