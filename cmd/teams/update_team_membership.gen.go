package teams

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateTeamMembershipCmd = &cobra.Command{
	Use:   "update-team-membership [team_id] [user_id]",
	Aliases: []string{ "update-membership", },
	Short: "Update a user's membership attributes on a team",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.UpdateTeamMembership(client.NewContext(apiKey, appKey, site), args[0], args[1], datadogV2.UserTeamUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-team-membership")

		cmdutil.PrintJSON(res, "team_memberships")
	},
}

func init() {
	Cmd.AddCommand(UpdateTeamMembershipCmd)
}
