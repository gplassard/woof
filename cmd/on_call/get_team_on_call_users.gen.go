package on_call

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetTeamOnCallUsersCmd = &cobra.Command{
	Use:   "get-team-on-call-users [team_id]",
	Aliases: []string{ "get-team-users", },
	Short: "Get team on-call users",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.GetTeamOnCallUsers(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-team-on-call-users")

		cmdutil.PrintJSON(res, "team_oncall_responders")
	},
}

func init() {
	Cmd.AddCommand(GetTeamOnCallUsersCmd)
}
