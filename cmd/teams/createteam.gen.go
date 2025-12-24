package teams

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateTeamCmd = &cobra.Command{
	Use:   "createteam",
	Short: "Create a team",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.CreateTeam(client.NewContext(apiKey, appKey, site), datadogV2.TeamCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createteam: %v", err)
		}

		cmdutil.PrintJSON(res, "teams")
	},
}

func init() {
	Cmd.AddCommand(CreateTeamCmd)
}
