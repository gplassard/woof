package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateTeamMembershipCmd = &cobra.Command{
	Use:     "update-team-membership [team_id] [user_id]",
	Aliases: []string{"update-membership"},
	Short:   "Update a user's membership attributes on a team",
	Long: `Update a user's membership attributes on a team
Documentation: https://docs.datadoghq.com/api/latest/teams/#update-team-membership`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UserTeamResponse
		var err error

		var body datadogV2.UserTeamUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err = api.UpdateTeamMembership(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-team-membership")

		cmd.Println(cmdutil.FormatJSON(res, "team_memberships"))
	},
}

func init() {

	UpdateTeamMembershipCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateTeamMembershipCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateTeamMembershipCmd)
}
