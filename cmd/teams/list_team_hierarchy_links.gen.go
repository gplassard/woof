package teams

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListTeamHierarchyLinksCmd = &cobra.Command{
	Use:   "list-team-hierarchy-links",
	Aliases: []string{ "list-hierarchy-links", },
	Short: "Get team hierarchy links",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.ListTeamHierarchyLinks(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-team-hierarchy-links")

		cmdutil.PrintJSON(res, "team_hierarchy_links")
	},
}

func init() {
	Cmd.AddCommand(ListTeamHierarchyLinksCmd)
}
