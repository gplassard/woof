package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateTeamMembershipCmd = &cobra.Command{
	Use:     "update-team-membership [team_id] [user_id] [payload]",
	Aliases: []string{"update-membership"},
	Short:   "Update a user's membership attributes on a team",
	Args:    cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.UserTeamUpdateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.UpdateTeamMembership(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-team-membership")

		cmd.Println(cmdutil.FormatJSON(res, "team_memberships"))
	},
}

func init() {
	Cmd.AddCommand(UpdateTeamMembershipCmd)
}
