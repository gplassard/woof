package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetTeamHierarchyLinkCmd = &cobra.Command{
	Use:   "getteamhierarchylink",
	Short: "Get a team hierarchy link",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/team-hierarchy-links/{link_id}")
		fmt.Println("OperationID: GetTeamHierarchyLink")
	},
}

func init() {
	Cmd.AddCommand(GetTeamHierarchyLinkCmd)
}
