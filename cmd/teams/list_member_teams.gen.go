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

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListMemberTeams(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-member-teams")

		fmt.Println(cmdutil.FormatJSON(res, "team"))
	},
}

func init() {

	Cmd.AddCommand(ListMemberTeamsCmd)
}
