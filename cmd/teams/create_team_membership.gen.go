package teams

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateTeamMembershipCmd = &cobra.Command{
	Use:   "create_team_membership [team_id]",
	Short: "Add a user to a team",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.CreateTeamMembership(client.NewContext(apiKey, appKey, site), args[0], datadogV2.UserTeamRequest{})
		if err != nil {
			log.Fatalf("failed to create_team_membership: %v", err)
		}

		cmdutil.PrintJSON(res, "team_memberships")
	},
}

func init() {
	Cmd.AddCommand(CreateTeamMembershipCmd)
}
