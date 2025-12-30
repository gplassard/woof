package on_call

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SetOnCallTeamRoutingRulesCmd = &cobra.Command{
	Use:     "set-on-call-team-routing-rules [team_id]",
	Aliases: []string{"set-team-routing-rules"},
	Short:   "Set On-Call team routing rules",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.SetOnCallTeamRoutingRules(client.NewContext(apiKey, appKey, site), args[0], datadogV2.TeamRoutingRulesRequest{})
		cmdutil.HandleError(err, "failed to set-on-call-team-routing-rules")

		cmd.Println(cmdutil.FormatJSON(res, "team_routing_rules"))
	},
}

func init() {
	Cmd.AddCommand(SetOnCallTeamRoutingRulesCmd)
}
