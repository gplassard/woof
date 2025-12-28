package teams

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var AddMemberTeamCmd = &cobra.Command{
	Use:   "add-member-team [super_team_id]",
	Aliases: []string{ "add-member", },
	Short: "Add a member team",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		_, err := api.AddMemberTeam(client.NewContext(apiKey, appKey, site), args[0], datadogV2.AddMemberTeamRequest{})
		cmdutil.HandleError(err, "failed to add-member-team")

		
	},
}

func init() {
	Cmd.AddCommand(AddMemberTeamCmd)
}
