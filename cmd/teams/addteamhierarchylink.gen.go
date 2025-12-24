package teams

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var AddTeamHierarchyLinkCmd = &cobra.Command{
	Use:   "addteamhierarchylink",
	Short: "Create a team hierarchy link",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.AddTeamHierarchyLink(client.NewContext(apiKey, appKey, site), datadogV2.TeamHierarchyLinkCreateRequest{})
		if err != nil {
			log.Fatalf("failed to addteamhierarchylink: %v", err)
		}

		cmdutil.PrintJSON(res, "teams")
	},
}

func init() {
	Cmd.AddCommand(AddTeamHierarchyLinkCmd)
}
