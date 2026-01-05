package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var AddMemberTeamCmd = &cobra.Command{
	Use:     "add-member-team [super_team_id] [payload]",
	Aliases: []string{"add-member"},
	Short:   "Add a member team",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.AddMemberTeamRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		_, err = api.AddMemberTeam(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to add-member-team")

	},
}

func init() {
	Cmd.AddCommand(AddMemberTeamCmd)
}
