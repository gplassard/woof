package teams

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var RemoveTeamHierarchyLinkCmd = &cobra.Command{
	Use:   "removeteamhierarchylink [link_id]",
	Short: "Remove a team hierarchy link",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		_, err := api.RemoveTeamHierarchyLink(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to removeteamhierarchylink: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(RemoveTeamHierarchyLinkCmd)
}
