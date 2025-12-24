package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var AddTeamHierarchyLinkCmd = &cobra.Command{
	Use:   "addteamhierarchylink",
	Short: "Create a team hierarchy link",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/team-hierarchy-links")
		fmt.Println("OperationID: AddTeamHierarchyLink")
	},
}

func init() {
	Cmd.AddCommand(AddTeamHierarchyLinkCmd)
}
