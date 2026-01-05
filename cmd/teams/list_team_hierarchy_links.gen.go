package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListTeamHierarchyLinksCmd = &cobra.Command{
	Use:     "list-team-hierarchy-links",
	Aliases: []string{"list-hierarchy-links"},
	Short:   "Get team hierarchy links",
	Long: `Get team hierarchy links
Documentation: https://docs.datadoghq.com/api/latest/teams/#list-team-hierarchy-links`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamHierarchyLinksResponse
		var err error

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err = api.ListTeamHierarchyLinks(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-team-hierarchy-links")

		cmd.Println(cmdutil.FormatJSON(res, "team_hierarchy_links"))
	},
}

func init() {

	Cmd.AddCommand(ListTeamHierarchyLinksCmd)
}
