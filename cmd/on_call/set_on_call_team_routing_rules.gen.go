package on_call

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var SetOnCallTeamRoutingRulesCmd = &cobra.Command{
	Use:     "set-on-call-team-routing-rules [team_id] [payload]",
	Aliases: []string{"set-team-routing-rules"},
	Short:   "Set On-Call team routing rules",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.TeamRoutingRulesRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.SetOnCallTeamRoutingRules(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to set-on-call-team-routing-rules")

		cmd.Println(cmdutil.FormatJSON(res, "team_routing_rules"))
	},
}

func init() {
	Cmd.AddCommand(SetOnCallTeamRoutingRulesCmd)
}
