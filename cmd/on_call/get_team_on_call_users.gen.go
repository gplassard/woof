package on_call

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetTeamOnCallUsersCmd = &cobra.Command{
	Use:     "get-team-on-call-users [team_id]",
	Aliases: []string{"get-team-users"},
	Short:   "Get team on-call users",
	Long: `Get team on-call users
Documentation: https://docs.datadoghq.com/api/latest/on-call/#get-team-on-call-users`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamOnCallResponders
		var err error

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err = api.GetTeamOnCallUsers(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-team-on-call-users")

		cmd.Println(cmdutil.FormatJSON(res, "team_oncall_responders"))
	},
}

func init() {

	Cmd.AddCommand(GetTeamOnCallUsersCmd)
}
