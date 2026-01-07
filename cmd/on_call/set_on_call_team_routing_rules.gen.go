package on_call

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SetOnCallTeamRoutingRulesCmd = &cobra.Command{
	Use:     "set-on-call-team-routing-rules [team_id]",
	Aliases: []string{"set-team-routing-rules"},
	Short:   "Set On-Call team routing rules",
	Long: `Set On-Call team routing rules
Documentation: https://docs.datadoghq.com/api/latest/on-call/#set-on-call-team-routing-rules`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamRoutingRules
		var err error

		optionalParams := datadogV2.NewSetOnCallTeamRoutingRulesOptionalParameters()

		if cmd.Flags().Changed("payload") || cmd.Flags().Changed("payload-file") {
			err = cmdutil.UnmarshalPayload(cmd, optionalParams)
			cmdutil.HandleError(err, "failed to read payload")
		}

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.SetOnCallTeamRoutingRules(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to set-on-call-team-routing-rules")

		cmd.Println(cmdutil.FormatJSON(res, "team_routing_rules"))
	},
}

func init() {

	SetOnCallTeamRoutingRulesCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	SetOnCallTeamRoutingRulesCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	SetOnCallTeamRoutingRulesCmd.Flags().String("include", "", "Comma-separated list of included relationships to be returned. Allowed values: 'rules', 'rules.policy'.")

	Cmd.AddCommand(SetOnCallTeamRoutingRulesCmd)
}
