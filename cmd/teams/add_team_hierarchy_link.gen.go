package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var AddTeamHierarchyLinkCmd = &cobra.Command{
	Use:     "add-team-hierarchy-link",
	Aliases: []string{"add-hierarchy-link"},
	Short:   "Create a team hierarchy link",
	Long: `Create a team hierarchy link
Documentation: https://docs.datadoghq.com/api/latest/teams/#add-team-hierarchy-link`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamHierarchyLinkResponse
		var err error

		var body datadogV2.TeamHierarchyLinkCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.AddTeamHierarchyLink(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to add-team-hierarchy-link")

		cmd.Println(cmdutil.FormatJSON(res, "team_hierarchy_links"))
	},
}

func init() {

	AddTeamHierarchyLinkCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	AddTeamHierarchyLinkCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(AddTeamHierarchyLinkCmd)
}
