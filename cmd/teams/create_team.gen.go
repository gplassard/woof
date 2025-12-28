package teams

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateTeamCmd = &cobra.Command{
	Use:   "create-team",
	Aliases: []string{ "create", },
	Short: "Create a team",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.CreateTeam(client.NewContext(apiKey, appKey, site), datadogV2.TeamCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create-team: %v", err)
		}

		cmdutil.PrintJSON(res, "team")
	},
}

func init() {
	Cmd.AddCommand(CreateTeamCmd)
}
