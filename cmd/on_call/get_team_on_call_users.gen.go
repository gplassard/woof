package on_call

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetTeamOnCallUsersCmd = &cobra.Command{
	Use:     "get-team-on-call-users [team_id]",
	Aliases: []string{"get-team-users"},
	Short:   "Get team on-call users",
	Long: `Get team on-call users
Documentation: https://docs.datadoghq.com/api/latest/on-call/#get-team-on-call-users`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamOnCallResponders
		var err error

		optionalParams := datadogV2.NewGetTeamOnCallUsersOptionalParameters()

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetTeamOnCallUsers(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to get-team-on-call-users")

		fmt.Println(cmdutil.FormatJSON(res, "team_oncall_responders"))
	},
}

func init() {

	GetTeamOnCallUsersCmd.Flags().String("include", "", "Comma-separated list of included relationships to be returned. Allowed values: 'responders', 'escalations', 'escalations.responders'.")

	Cmd.AddCommand(GetTeamOnCallUsersCmd)
}
