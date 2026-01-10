package teams

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListMemberTeamsCmd = &cobra.Command{
	Use:     "list-member-teams [super_team_id]",
	Aliases: []string{"list-member"},
	Short:   "Get all member teams",
	Long: `Get all member teams
Documentation: https://docs.datadoghq.com/api/latest/teams/#list-member-teams`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamsResponse
		var err error

		optionalParams := datadogV2.NewListMemberTeamsOptionalParameters()

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("page-number") {
			val, _ := cmd.Flags().GetInt64("page-number")
			optionalParams.WithPageNumber(val)
		}

		if cmd.Flags().Changed("fields-team") {
			val, _ := cmd.Flags().GetString("fields-team")
			optionalParams.WithFieldsTeam(val)
		}

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListMemberTeams(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to list-member-teams")

		fmt.Println(cmdutil.FormatJSON(res, "team"))
	},
}

func init() {

	ListMemberTeamsCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	ListMemberTeamsCmd.Flags().Int64("page-number", 0, "Specific page number to return.")

	ListMemberTeamsCmd.Flags().String("fields-team", "", "List of fields that need to be fetched.")

	Cmd.AddCommand(ListMemberTeamsCmd)
}
