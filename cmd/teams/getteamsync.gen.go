package teams

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetTeamSyncCmd = &cobra.Command{
	Use:   "getteamsync [filter[source]]",
	Short: "Get team sync configurations",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.GetTeamSync(client.NewContext(apiKey, appKey, site), datadogV2.TeamSyncAttributesSource(args[0]))
		if err != nil {
			log.Fatalf("failed to getteamsync: %v", err)
		}

		cmdutil.PrintJSON(res, "teams")
	},
}

func init() {
	Cmd.AddCommand(GetTeamSyncCmd)
}
