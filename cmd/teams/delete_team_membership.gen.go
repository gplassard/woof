package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteTeamMembershipCmd = &cobra.Command{
	Use:     "delete-team-membership [team_id] [user_id]",
	Aliases: []string{"delete-membership"},
	Short:   "Remove a user from a team",
	Long: `Remove a user from a team
Documentation: https://docs.datadoghq.com/api/latest/teams/#delete-team-membership`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteTeamMembership(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to delete-team-membership")

	},
}

func init() {

	Cmd.AddCommand(DeleteTeamMembershipCmd)
}
