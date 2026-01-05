package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetUserMembershipsCmd = &cobra.Command{
	Use: "get-user-memberships [user_uuid]",

	Short: "Get user memberships",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UserTeamsResponse
		var err error

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err = api.GetUserMemberships(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-user-memberships")

		cmd.Println(cmdutil.FormatJSON(res, "team_memberships"))
	},
}

func init() {
	Cmd.AddCommand(GetUserMembershipsCmd)
}
