package on_call

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetTeamOnCallUsersCmd = &cobra.Command{
	Use:   "getteamoncallusers [team_id]",
	Short: "Get team on-call users",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.GetTeamOnCallUsers(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getteamoncallusers: %v", err)
		}

		cmdutil.PrintJSON(res, "on_call")
	},
}

func init() {
	Cmd.AddCommand(GetTeamOnCallUsersCmd)
}
