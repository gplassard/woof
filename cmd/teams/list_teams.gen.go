package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListTeamsCmd = &cobra.Command{
	Use:     "list-teams",
	Aliases: []string{"list"},
	Short:   "Get all teams",
	Long: `Get all teams
Documentation: https://docs.datadoghq.com/api/latest/teams/#list-teams`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamsResponse
		var err error

		optionalParams := datadogV2.NewListTeamsOptionalParameters()

		if cmd.Flags().Changed("page-number") {
			val, _ := cmd.Flags().GetInt64("page-number")
			optionalParams.WithPageNumber(val)
		}

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("sort") {
			val, _ := cmd.Flags().GetString("sort")
			optionalParams.WithSort(val)
		}

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		if cmd.Flags().Changed("filter-keyword") {
			val, _ := cmd.Flags().GetString("filter-keyword")
			optionalParams.WithFilterKeyword(val)
		}

		if cmd.Flags().Changed("filter-me") {
			val, _ := cmd.Flags().GetString("filter-me")
			optionalParams.WithFilterMe(val)
		}

		if cmd.Flags().Changed("fields-team") {
			val, _ := cmd.Flags().GetString("fields-team")
			optionalParams.WithFieldsTeam(val)
		}

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListTeams(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-teams")

		cmd.Println(cmdutil.FormatJSON(res, "team"))
	},
}

func init() {

	ListTeamsCmd.Flags().Int64("page-number", 0, "Specific page number to return.")

	ListTeamsCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	ListTeamsCmd.Flags().String("sort", "", "Specifies the order of the returned teams")

	ListTeamsCmd.Flags().String("include", "", "Included related resources optionally requested. Allowed enum values: 'team_links, user_team_permissions'")

	ListTeamsCmd.Flags().String("filter-keyword", "", "Search query. Can be team name, team handle, or email of team member")

	ListTeamsCmd.Flags().String("filter-me", "", "When true, only returns teams the current user belongs to")

	ListTeamsCmd.Flags().String("fields-team", "", "List of fields that need to be fetched.")

	Cmd.AddCommand(ListTeamsCmd)
}
