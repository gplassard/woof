package teams

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetUserMembershipsCmd = &cobra.Command{
	Use:   "get-user-memberships [user_uuid]",
	
	Short: "Get user memberships",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.GetUserMemberships(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get-user-memberships: %v", err)
		}

		cmdutil.PrintJSON(res, "team_memberships")
	},
}

func init() {
	Cmd.AddCommand(GetUserMembershipsCmd)
}
