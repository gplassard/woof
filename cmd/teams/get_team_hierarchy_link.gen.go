package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetTeamHierarchyLinkCmd = &cobra.Command{
	Use:     "get-team-hierarchy-link [link_id]",
	Aliases: []string{"get-hierarchy-link"},
	Short:   "Get a team hierarchy link",
	Long: `Get a team hierarchy link
Documentation: https://docs.datadoghq.com/api/latest/teams/#get-team-hierarchy-link`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamHierarchyLinkResponse
		var err error

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetTeamHierarchyLink(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-team-hierarchy-link")

		cmd.Println(cmdutil.FormatJSON(res, "team_hierarchy_links"))
	},
}

func init() {

	Cmd.AddCommand(GetTeamHierarchyLinkCmd)
}
