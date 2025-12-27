package teams

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateTeamPermissionSettingCmd = &cobra.Command{
	Use:   "updateteampermissionsetting [team_id] [action]",
	Short: "Update permission setting for team",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.UpdateTeamPermissionSetting(client.NewContext(apiKey, appKey, site), args[0], args[1], datadogV2.TeamPermissionSettingUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to updateteampermissionsetting: %v", err)
		}

		cmdutil.PrintJSON(res, "teams")
	},
}

func init() {
	Cmd.AddCommand(UpdateTeamPermissionSettingCmd)
}
