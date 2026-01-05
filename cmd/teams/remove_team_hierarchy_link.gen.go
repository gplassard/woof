package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var RemoveTeamHierarchyLinkCmd = &cobra.Command{
	Use:     "remove-team-hierarchy-link [link_id]",
	Aliases: []string{"remove-hierarchy-link"},
	Short:   "Remove a team hierarchy link",
	Long: `Remove a team hierarchy link
Documentation: https://docs.datadoghq.com/api/latest/teams/#remove-team-hierarchy-link`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		_, err = api.RemoveTeamHierarchyLink(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to remove-team-hierarchy-link")

	},
}

func init() {

	Cmd.AddCommand(RemoveTeamHierarchyLinkCmd)
}
