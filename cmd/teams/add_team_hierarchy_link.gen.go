package teams

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var AddTeamHierarchyLinkCmd = &cobra.Command{
	Use:   "add-team-hierarchy-link",
	Aliases: []string{ "add-hierarchy-link", },
	Short: "Create a team hierarchy link",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.AddTeamHierarchyLink(client.NewContext(apiKey, appKey, site), datadogV2.TeamHierarchyLinkCreateRequest{})
		if err != nil {
			log.Fatalf("failed to add-team-hierarchy-link: %v", err)
		}

		cmdutil.PrintJSON(res, "team_hierarchy_links")
	},
}

func init() {
	Cmd.AddCommand(AddTeamHierarchyLinkCmd)
}
