package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteTeamCmd = &cobra.Command{
	Use:     "delete-team [team_id]",
	Aliases: []string{"delete"},
	Short:   "Remove a team",
	Long: `Remove a team
Documentation: https://docs.datadoghq.com/api/latest/teams/#delete-team`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		_, err = api.DeleteTeam(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-team")

	},
}

func init() {

	Cmd.AddCommand(DeleteTeamCmd)
}
