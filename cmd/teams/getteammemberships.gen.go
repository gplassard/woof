package teams

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetTeamMembershipsCmd = &cobra.Command{
	Use:   "getteammemberships [team_id]",
	Short: "Get team memberships",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.GetTeamMemberships(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getteammemberships: %v", err)
		}

		cmdutil.PrintJSON(res, "teams")
	},
}

func init() {
	Cmd.AddCommand(GetTeamMembershipsCmd)
}
