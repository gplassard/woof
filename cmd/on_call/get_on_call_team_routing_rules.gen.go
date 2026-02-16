package on_call

import (
	"fmt"
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

		optionalParams := datadogV2.NewGetOnCallTeamRoutingRulesOptionalParameters()

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetOnCallTeamRoutingRules(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to get-on-call-team-routing-rules")

		fmt.Println(cmdutil.FormatJSON(res, "team_routing_rules"))
	},
}

func init() {

	GetOnCallTeamRoutingRulesCmd.Flags().String("include", "", "Comma-separated list of included relationships to be returned. Allowed values: 'rules', 'rules.policy'.")

	Cmd.AddCommand(GetOnCallTeamRoutingRulesCmd)
}
