package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListTeamHierarchyLinksCmd = &cobra.Command{
	Use:   "listteamhierarchylinks",
	Short: "Get team hierarchy links",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/team-hierarchy-links")
		fmt.Println("OperationID: ListTeamHierarchyLinks")
	},
}

func init() {
	Cmd.AddCommand(ListTeamHierarchyLinksCmd)
}
