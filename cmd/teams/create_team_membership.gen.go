package teams

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateTeamMembershipCmd = &cobra.Command{
	Use:     "create-team-membership [team_id]",
	Aliases: []string{"create-membership"},
	Short:   "Add a user to a team",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.CreateTeamMembership(client.NewContext(apiKey, appKey, site), args[0], datadogV2.UserTeamRequest{})
		cmdutil.HandleError(err, "failed to create-team-membership")

		cmd.Println(cmdutil.FormatJSON(res, "team_memberships"))
	},
}

func init() {
	Cmd.AddCommand(CreateTeamMembershipCmd)
}
