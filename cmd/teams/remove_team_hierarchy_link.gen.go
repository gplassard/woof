package teams

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var RemoveTeamHierarchyLinkCmd = &cobra.Command{
	Use:   "remove-team-hierarchy-link [link_id]",
	Aliases: []string{ "remove-hierarchy-link", },
	Short: "Remove a team hierarchy link",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		_, err := api.RemoveTeamHierarchyLink(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to remove-team-hierarchy-link: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(RemoveTeamHierarchyLinkCmd)
}
