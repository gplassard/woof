package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var AddTeamHierarchyLinkCmd = &cobra.Command{
	Use:     "add-team-hierarchy-link [payload]",
	Aliases: []string{"add-hierarchy-link"},
	Short:   "Create a team hierarchy link",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.TeamHierarchyLinkCreateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.AddTeamHierarchyLink(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to add-team-hierarchy-link")

		cmd.Println(cmdutil.FormatJSON(res, "team_hierarchy_links"))
	},
}

func init() {
	Cmd.AddCommand(AddTeamHierarchyLinkCmd)
}
