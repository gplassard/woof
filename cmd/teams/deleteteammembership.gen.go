package teams

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteTeamMembershipCmd = &cobra.Command{
	Use:   "deleteteammembership [team_id] [user_id]",
	Short: "Remove a user from a team",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		_, err := api.DeleteTeamMembership(client.NewContext(apiKey, appKey, site), args[0], args[1])
		if err != nil {
			log.Fatalf("failed to deleteteammembership: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteTeamMembershipCmd)
}
