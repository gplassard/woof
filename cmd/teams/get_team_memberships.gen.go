package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetTeamMembershipsCmd = &cobra.Command{
	Use:     "get-team-memberships [team_id]",
	Aliases: []string{"get-memberships"},
	Short:   "Get team memberships",
	Long: `Get team memberships
Documentation: https://docs.datadoghq.com/api/latest/teams/#get-team-memberships`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UserTeamsResponse
		var err error

		optionalParams := datadogV2.NewGetTeamMembershipsOptionalParameters()

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("page-number") {
			val, _ := cmd.Flags().GetInt64("page-number")
			optionalParams.WithPageNumber(val)
		}

		if cmd.Flags().Changed("sort") {
			val, _ := cmd.Flags().GetString("sort")
			optionalParams.WithSort(val)
		}

		if cmd.Flags().Changed("filter-keyword") {
			val, _ := cmd.Flags().GetString("filter-keyword")
			optionalParams.WithFilterKeyword(val)
		}

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetTeamMemberships(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to get-team-memberships")

		cmd.Println(cmdutil.FormatJSON(res, "team_memberships"))
	},
}

func init() {

	GetTeamMembershipsCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	GetTeamMembershipsCmd.Flags().Int64("page-number", 0, "Specific page number to return.")

	GetTeamMembershipsCmd.Flags().String("sort", "", "Specifies the order of returned team memberships")

	GetTeamMembershipsCmd.Flags().String("filter-keyword", "", "Search query, can be user email or name")

	Cmd.AddCommand(GetTeamMembershipsCmd)
}
