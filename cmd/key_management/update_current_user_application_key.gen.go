package key_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateCurrentUserApplicationKeyCmd = &cobra.Command{
	Use:   "update_current_user_application_key [app_key_id]",
	Short: "Edit an application key owned by current user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		res, _, err := api.UpdateCurrentUserApplicationKey(client.NewContext(apiKey, appKey, site), args[0], datadogV2.ApplicationKeyUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to update_current_user_application_key: %v", err)
		}

		cmdutil.PrintJSON(res, "key_management")
	},
}

func init() {
	Cmd.AddCommand(UpdateCurrentUserApplicationKeyCmd)
}
