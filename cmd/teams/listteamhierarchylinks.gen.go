package teams

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListTeamHierarchyLinksCmd = &cobra.Command{
	Use:   "listteamhierarchylinks",
	Short: "Get team hierarchy links",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.ListTeamHierarchyLinks(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listteamhierarchylinks: %v", err)
		}

		cmdutil.PrintJSON(res, "teams")
	},
}

func init() {
	Cmd.AddCommand(ListTeamHierarchyLinksCmd)
}
