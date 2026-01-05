package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateTeamMembershipCmd = &cobra.Command{
	Use:     "create-team-membership [team_id]",
	Aliases: []string{"create-membership"},
	Short:   "Add a user to a team",
	Long: `Add a user to a team
Documentation: https://docs.datadoghq.com/api/latest/teams/#create-team-membership`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UserTeamResponse
		var err error

		var body datadogV2.UserTeamRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err = api.CreateTeamMembership(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-team-membership")

		cmd.Println(cmdutil.FormatJSON(res, "team_memberships"))
	},
}

func init() {

	CreateTeamMembershipCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateTeamMembershipCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateTeamMembershipCmd)
}
