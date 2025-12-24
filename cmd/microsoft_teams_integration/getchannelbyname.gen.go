package microsoft_teams_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetChannelByNameCmd = &cobra.Command{
	Use:   "getchannelbyname",
	Short: "Get channel information by name",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integration/ms-teams/configuration/channel/{tenant_name}/{team_name}/{channel_name}")
		fmt.Println("OperationID: GetChannelByName")
	},
}

func init() {
	Cmd.AddCommand(GetChannelByNameCmd)
}
