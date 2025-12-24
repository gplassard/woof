package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RemoveTeamHierarchyLinkCmd = &cobra.Command{
	Use:   "removeteamhierarchylink",
	Short: "Remove a team hierarchy link",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/team-hierarchy-links/{link_id}")
		fmt.Println("OperationID: RemoveTeamHierarchyLink")
	},
}

func init() {
	Cmd.AddCommand(RemoveTeamHierarchyLinkCmd)
}
