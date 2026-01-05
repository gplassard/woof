package on_call

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetOnCallTeamRoutingRulesCmd = &cobra.Command{
	Use:     "get-on-call-team-routing-rules [team_id]",
	Aliases: []string{"get-team-routing-rules"},
	Short:   "Get On-Call team routing rules",
	Long: `Get On-Call team routing rules
Documentation: https://docs.datadoghq.com/api/latest/on-call/#get-on-call-team-routing-rules`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamRoutingRules
		var err error

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err = api.GetOnCallTeamRoutingRules(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-on-call-team-routing-rules")

		cmd.Println(cmdutil.FormatJSON(res, "team_routing_rules"))
	},
}

func init() {

	Cmd.AddCommand(GetOnCallTeamRoutingRulesCmd)
}
