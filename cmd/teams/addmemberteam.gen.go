package teams

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var AddMemberTeamCmd = &cobra.Command{
	Use:   "addmemberteam [super_team_id]",
	Short: "Add a member team",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		_, err := api.AddMemberTeam(client.NewContext(apiKey, appKey, site), args[0], datadogV2.AddMemberTeamRequest{})
		if err != nil {
			log.Fatalf("failed to addmemberteam: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(AddMemberTeamCmd)
}
