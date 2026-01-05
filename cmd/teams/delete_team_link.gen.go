package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteTeamLinkCmd = &cobra.Command{
	Use:     "delete-team-link [team_id] [link_id]",
	Aliases: []string{"delete-link"},
	Short:   "Remove a team link",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		_, err := api.DeleteTeamLink(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to delete-team-link")

	},
}

func init() {
	Cmd.AddCommand(DeleteTeamLinkCmd)
}
