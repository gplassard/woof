package teams

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetTeamHierarchyLinkCmd = &cobra.Command{
	Use:   "getteamhierarchylink [link_id]",
	Short: "Get a team hierarchy link",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.GetTeamHierarchyLink(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getteamhierarchylink: %v", err)
		}

		cmdutil.PrintJSON(res, "teams")
	},
}

func init() {
	Cmd.AddCommand(GetTeamHierarchyLinkCmd)
}
