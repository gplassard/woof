package teams

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetTeamHierarchyLinkCmd = &cobra.Command{
	Use:   "get-team-hierarchy-link [link_id]",
	Aliases: []string{ "get-hierarchy-link", },
	Short: "Get a team hierarchy link",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.GetTeamHierarchyLink(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get-team-hierarchy-link: %v", err)
		}

		cmdutil.PrintJSON(res, "team_hierarchy_links")
	},
}

func init() {
	Cmd.AddCommand(GetTeamHierarchyLinkCmd)
}
