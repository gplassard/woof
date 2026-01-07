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

		optionalParams := datadogV2.NewListTeamHierarchyLinksOptionalParameters()

		if cmd.Flags().Changed("page-number") {
			val, _ := cmd.Flags().GetInt64("page-number")
			optionalParams.WithPageNumber(val)
		}

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("filter-parent-team") {
			val, _ := cmd.Flags().GetString("filter-parent-team")
			optionalParams.WithFilterParentTeam(val)
		}

		if cmd.Flags().Changed("filter-sub-team") {
			val, _ := cmd.Flags().GetString("filter-sub-team")
			optionalParams.WithFilterSubTeam(val)
		}

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListTeamHierarchyLinks(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-team-hierarchy-links")

		cmd.Println(cmdutil.FormatJSON(res, "team_hierarchy_links"))
	},
}

func init() {

	ListTeamHierarchyLinksCmd.Flags().Int64("page-number", 0, "Specific page number to return.")

	ListTeamHierarchyLinksCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	ListTeamHierarchyLinksCmd.Flags().String("filter-parent-team", "", "Filter by parent team ID")

	ListTeamHierarchyLinksCmd.Flags().String("filter-sub-team", "", "Filter by sub team ID")

	Cmd.AddCommand(ListTeamHierarchyLinksCmd)
}
