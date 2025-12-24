package teams

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateTeamConnectionsCmd = &cobra.Command{
	Use:   "createteamconnections",
	Short: "Create team connections",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.CreateTeamConnections(client.NewContext(apiKey, appKey, site), datadogV2.TeamConnectionCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createteamconnections: %v", err)
		}

		cmdutil.PrintJSON(res, "teams")
	},
}

func init() {
	Cmd.AddCommand(CreateTeamConnectionsCmd)
}
