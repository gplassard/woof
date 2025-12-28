package teams

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListMemberTeamsCmd = &cobra.Command{
	Use:   "list-member-teams [super_team_id]",
	Aliases: []string{ "list-member", },
	Short: "Get all member teams",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.ListMemberTeams(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to list-member-teams: %v", err)
		}

		cmdutil.PrintJSON(res, "team")
	},
}

func init() {
	Cmd.AddCommand(ListMemberTeamsCmd)
}
