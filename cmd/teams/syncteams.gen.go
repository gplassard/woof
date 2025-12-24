package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var SyncTeamsCmd = &cobra.Command{
	Use:   "syncteams",
	Short: "Link Teams with GitHub Teams",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/team/sync")
		fmt.Println("OperationID: SyncTeams")
	},
}

func init() {
	Cmd.AddCommand(SyncTeamsCmd)
}
